import React, { useState, useEffect } from "react";
import axios from "axios";
import { Button, Form, Container, Modal } from "react-bootstrap";
import Entry from "./Entry";

const Entries = () => {
  const [entries, setEntries] = useState([]);
  const [refreshData, setRefreshData] = useState(false);
  const [changeEntry, setChangeEntry] = useState({ change: false, id: 0 });
  const [changeIngredient, setChangeIngredient] = useState({
    change: false,
    id: 0,
  });
  const [newIngredientName, setNewIngredientName] = useState("");
  const [addNewEntry, setAddNewEntry] = useState(false);
  const [newEntry, setNewEntry] = useState({
    dish: "",
    ingredients: "",
    calories: 0,
    fat: 0,
  });

  useEffect(() => {
    getAllEntries();
  }, []);

  useEffect(() => {
    if (refreshData) {
      setRefreshData(false);
      getAllEntries();
    }
  }, [refreshData]);

  return (
    <div>
      <Container>
        <Button onClick={() => setAddNewEntry(true)}>
          Track today's calories
        </Button>
      </Container>
      <Container>
        {entries &&
          entries.map((entry, i) => (
            <Entry
              key={i}
              entryData={entry}
              deleteSingleEntry={deleteSingleEntry}
              setChangeIngredient={setChangeIngredient}
              setChangeEntry={setChangeEntry}
            />
          ))}
      </Container>

      <Modal
        show={addNewEntry}
        onHide={() => setAddNewEntry(false)}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Add Calorie Entry</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Dish</Form.Label>
            <Form.Control
              onChange={(event) => {
                newEntry.dish = event.target.value;
              }}
            />
            <Form.Label>Ingredients</Form.Label>
            <Form.Control
              onChange={(event) => {
                newEntry.ingredients = event.target.value;
              }}
            />
            <Form.Label>Calories</Form.Label>
            <Form.Control
              onChange={(event) => {
                newEntry.calories = event.target.value;
              }}
            />
            <Form.Label>Fat</Form.Label>
            <Form.Control
              type="number"
              onChange={(event) => {
                newEntry.fat = event.target.value;
              }}
            />
          </Form.Group>
          <Button onClick={() => addSingleEntry()}>Add</Button>
          <Button onClick={() => setAddNewEntry(false)}>Cancel</Button>
        </Modal.Body>
      </Modal>

      <Modal
        show={changeIngredient.change}
        onHide={() => setChangeIngredient({ change: false, id: 0 })}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Change Ingredients</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>New Ingredients</Form.Label>
            <Form.Control
              onChange={(event) => {
                setNewIngredientName(event.target.value);
              }}
            />
          </Form.Group>
          <Button onClick={() => changeIngredientForEntry()}>Change</Button>
          <Button
            onClick={() => setChangeIngredient({ change: false, id: 0 })}
          >
            Cancel
          </Button>
        </Modal.Body>
      </Modal>

      <Modal
        show={changeEntry.change}
        onHide={() => setChangeEntry({ change: false, id: 0 })}
        centered
      >
        <Modal.Header closeButton>
          <Modal.Title>Change Entry</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group>
            <Form.Label>Dish</Form.Label>
            <Form.Control
              onChange={(event) => {
                newEntry.dish = event.target.value;
              }}
            />
            <Form.Label>Ingredients</Form.Label>
            <Form.Control
              onChange={(event) => {
                newEntry.ingredients = event.target.value;
              }}
            />
            <Form.Label>Calories</Form.Label>
            <Form.Control
              onChange={(event) => {
                newEntry.calories = event.target.value;
              }}
            />
            <Form.Label>Fat</Form.Label>
            <Form.Control
              type="number"
              onChange={(event) => {
                newEntry.fat = event.target.value;
              }}
            />
          </Form.Group>
          <Button onClick={() => changeSingleEntry()}>Change</Button>
          <Button onClick={() => setChangeEntry({ change: false, id: 0 })}>
            Cancel
          </Button>
        </Modal.Body>
      </Modal>
    </div>
  );

  function changeIngredientForEntry() {
    const url =
      "http://localhost:8000/ingredient/update/" + changeIngredient.id;
    axios.put(url, { ingredients: newIngredientName }).then((response) => {
      if (response.status === 200) {
        setRefreshData(true);
      }
    });
    setChangeIngredient({ change: false, id: 0 });
  }

  function changeSingleEntry() {
    const url = "http://localhost:8000/entry/update/" + changeEntry.id;
    axios.put(url, newEntry).then((response) => {
      if (response.status === 200) {
        setRefreshData(true);
      }
    });
    setChangeEntry({ change: false, id: 0 });
  }

  function addSingleEntry() {
    const url = "http://localhost:8000/entry/create";
    axios
      .post(url, {
        dish: newEntry.dish,
        ingredients: newEntry.ingredients,
        calories: newEntry.calories,
        fat: parseFloat(newEntry.fat),
      })
      .then((response) => {
        if (response.status === 200) {
          setRefreshData(true);
        }
      });
    setAddNewEntry(false);
  }

  function deleteSingleEntry(id) {
    const url = "http://localhost:8000/entry/delete/" + id;
    axios.delete(url).then((response) => {
      if (response.status === 200) {
        setRefreshData(true);
      }
    });
  }

  function getAllEntries() {
    const url = "http://localhost:8000/entries";
    axios.get(url, { responseType: "json" }).then((response) => {
      if (response.status === 200) {
        setEntries(response.data);
      }
    });
  }
};

export default Entries;