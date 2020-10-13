const express = require('express')
const bodyParser = require('body-parser')
const app = express()
app.use(bodyParser.json({ type: 'application/*+json' }));

const port = 3000

app.get('/dapr/subscribe', (req, res) => {
    res.json([
        {
            pubsubname: "messagebus-node",
            topic: "songs",
            route: "playlist"
        }
    ]);
})

app.post('/playlist', (req, res) => {
    let song = req.body.data;
    console.log("New song request: " + song.artist + " - " + song.name);
    res.sendStatus(200);
});

app.listen(port, () => console.log(`consumer app listening on port ${port}!`))	
