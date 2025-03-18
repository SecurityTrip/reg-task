const express = require("express");
const axios = require("axios");

const app = express();
const port = process.env.PORT || 80;
const backendUrl = process.env.BACKEND_URL || "http://backend:8080";

app.use(express.json());

app.post("/register", async (req, res) => {
  try {
    console.log(`Проксируем на ${backendUrl}/register`, req.body);
    const response = await axios.post(`${backendUrl}/register`, req.body);
    res.json(response.data);
  } catch (error) {
    console.error("Ошибка проксирования:", error.message);
    res.status(error.response?.status || 500).json({ error: error.message });
  }
});

app.use(express.static("public"));

app.listen(port, () => {
  console.log(`Frontend работает на порту ${port}`);
});
