package routes

import (
    "context"
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "backend/models"
)

// Global variables
var validate = validator.New()
var entryCollection *mongo.Collection = OpenCollection("calories") // Use OpenConnection correctly
var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// AddEntry creates a new entry in the calories collection.
func AddEntry(c *gin.Context) {
    defer cancel() // Ensure to call cancel at the end of the function

    var entry models.Entry

    if err := c.BindJSON(&entry); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }
    validationErr := validate.Struct(entry)
    if validationErr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
        fmt.Println(validationErr)
        return
    }

    entry.ID = primitive.NewObjectID()
    result, insertErr := entryCollection.InsertOne(ctx, entry)
    if insertErr != nil {
        msg := fmt.Sprintf("Entry was not created")
        c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
        fmt.Println(insertErr)
        return
    }

    c.JSON(http.StatusOK, result)
}

// GetEntries retrieves all entries from the calories collection.
func GetEntries(c *gin.Context) {
    defer cancel()

    var entries []bson.M
    cursor, err := entryCollection.Find(ctx, bson.M{})

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    if err = cursor.All(ctx, &entries); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK, entries)
}

// GetEntriesByIngredient searches entries by a specific ingredient.
func GetEntriesByIngredient(c *gin.Context) {
    ingredient := c.Params.ByName("ingredient")
    defer cancel()

    var entries []bson.M
    cursor, err := entryCollection.Find(ctx, bson.M{"ingredients": ingredient})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    if err = cursor.All(ctx, &entries); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK, entries)
}

// GetEntryById retrieves a specific entry by its ID.
func GetEntryById(c *gin.Context) {
    EntryID := c.Params.ByName("id")
    docID, err := primitive.ObjectIDFromHex(EntryID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    defer cancel()
    var entry bson.M
    if err := entryCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&entry); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK, entry)
}

// UpdateIngredient updates the ingredients for a specific entry by ID.
func UpdateIngredient(c *gin.Context) {
    entryID := c.Params.ByName("id")
    docID, err := primitive.ObjectIDFromHex(entryID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    defer cancel()
    type Ingredient struct {
        Ingredients *string `json:"ingredients"`
    }
    var ingredient Ingredient

    if err := c.BindJSON(&ingredient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    result, err := entryCollection.UpdateOne(ctx, bson.M{"_id": docID},
        bson.D{{"$set", bson.D{{"ingredients", ingredient.Ingredients}}}}, 
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK, result.ModifiedCount)
}

// UpdateEntry updates an existing entry by its ID.
func UpdateEntry(c *gin.Context) {
    entryID := c.Params.ByName("id")
    docID, err := primitive.ObjectIDFromHex(entryID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    defer cancel()
    var entry models.Entry

    if err := c.BindJSON(&entry); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    validationErr := validate.Struct(entry)
    if validationErr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
        fmt.Println(validationErr)
        return
    }

    result, err := entryCollection.ReplaceOne(
        ctx,
        bson.M{"_id": docID},
        bson.M{
            "dish":        entry.Dish,
            "fat":         entry.Fat,
            "ingredients": entry.Ingredients,
            "calories":    entry.Calories,
        },
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK, result.ModifiedCount)
}

// DeleteEntry removes an entry by its ID.
func DeleteEntry(c *gin.Context) {
    entryID := c.Params.ByName("id")
    docID, err := primitive.ObjectIDFromHex(entryID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    defer cancel()
    result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK, result.DeletedCount)
}