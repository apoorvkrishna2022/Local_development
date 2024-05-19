import express, { request } from "express";
import { parse } from "nodemon/lib/cli/index.js";

const app = express();

/*
 * this is used in the case of post request where we want to send reqeust body or payload
 * */
app.use(express.json());

const PORT = process.env.PORT || 3000;

const mockUsers = [
  { id: 1, userName: "anson", displayName: "Anson" },
  { id: 2, userName: "jack", displayName: "Jack" },
  { id: 3, userName: "adam", displayName: "Adam" },
];

const mockProducts = [
  { id: 1, name: "chicken brest", price: 12.99 },
  { id: 2, name: "chicken leg", price: 11.99 },
];

app.listen(PORT, () => {
  console.log(`Running on Port ${PORT}`);
});

app.get("/", (request, response) => {
  response.status(200).send({ msg: "Hello World" });
});

//query params are ?key1=value&key2=value2
app.get("/api/users", (request, response) => {
  const {
    query: { filter = undefined, value = undefined },
  } = request;

  //when filters are undefined
  if (filter && value) {
    return response
      .status(200)
      .send(mockUsers.filter((user) => user[filter].includes(value)));
  }
  return response.status(200).send(mockUsers);
});

app.get("/api/products", (request, response) => {
  response.status(200).send(mockProducts);
});

//:id is the route parameter
/*
 * difference between status and sendStatus.
 * in status it just sets the status code and you can still send the response by send
 * in sendStatus it not only sent the status it dont allow you to send the response body. all it sends is ok as response.
 * */
app.get("/api/user/:id", (request, response) => {
  const parsedId = parseInt(request.params.id);
  if (isNaN(parsedId)) {
    return response.sendStatus(400);
  }
  const findUser = mockUsers.find((user) => user.id === parsedId);
  if (!findUser) {
    return response.status(200).send({ msg: "Couldn't find the user." });
  }
  response.status(200).send(findUser);
});

app.post("/api/user", (request, response) => {
  const { body: payload } = request;
  const newUser = { id: mockUsers[mockUsers.length - 1].id + 1, ...payload };
  mockUsers.push(newUser);
  return response.status(200).send(newUser);
});

/*
 * PUT and PATCH methods are used to update the data(record)
 * PATCH is used to partial update a record
 * PUT is used to update the entire record.
 * */

app.put("/api/user/:id", (request, response) => {
  const {
    body: payload,
    params: { id },
  } = request;

  const parsedId = parseInt(id);
  if (isNaN(parsedId)) {
    return response.sendStatus(400);
  }

  const findUserIndex = mockUsers.findIndex((user) => user.id === parsedId);

  if (findUserIndex === -1) {
    return response.status(200).send({ msg: "Couldn't find the User" });
  }

  mockUsers[findUserIndex] = { id: parsedId, ...payload };

  return response.status(200).send(mockUsers[findUserIndex]);
});

app.patch("/api/user/:id", (request, response) => {
  const {
    body: payload,
    params: { id },
  } = request;

  const parsedId = parseInt(id);
  if (isNaN(parsedId)) {
    return response.sendStatus(400);
  }

  const findUserIndex = mockUsers.findIndex((user) => user.id === parsedId);

  if (findUserIndex === -1) {
    return response.status(200).send({ msg: "Couldn't find the User" });
  }

  mockUsers[findUserIndex] = { ...mockUsers[findUserIndex], ...payload };

  return response.status(200).send(mockUsers[findUserIndex]);
});

app.delete("/api/user/:id", (request, response) => {
  const {
    params: { id },
  } = request;

  const parsedId = parseInt(id);
  if (isNaN(parsedId)) {
    return response.sendStatus(400);
  }

  const findUserIndex = mockUsers.findIndex((user) => user.id === parsedId);

  if (findUserIndex === -1) {
    return response.status(200).send({ msg: "Couldn't find the User" });
  }

  mockUsers.splice(findUserIndex, 1);

  return response.sendStatus(200);
});
