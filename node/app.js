const express = require('express')
const bodyParser = require('body-parser')
const app = express()
app.use(bodyParser.json())

const port = 3000

app.get('/dapr/subscribe', (req, res) => {
    res.json([
        {
            topic: "songs",
            route: "archive"
        }
    ]);
})

app.post('/archive', (req, res) => {
    console.log('worked!');
    res.sendStatus(200);
});

app.listen(port, () => console.log(`consumer app listening on port ${port}!`))	
