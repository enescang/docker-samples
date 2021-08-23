const express = require("express");

const app = express();

app.get('/', (req, res) => {
    res.json({ message: 'Hello Docker' });
})

const port = process.env.port || 3000;
app.listen(port, () => {
    console.log(`Server listening on ${port}...`);
});
