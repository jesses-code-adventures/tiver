# tiver

testing river to get my head around request queues

## goal

be able to hit a server with http requests that are received by a browser app, processed on the client and then responded to.

when a client receives a request:
- a square with a given width and colour is created on the top or left edge of the screen.
- the square starts to journey across the viewport at a fixed speed.
- if the square makes it all the way across the screen with no collisions, respond with a 200.
- if the square collides with another square, both squares should respond with 429s for the first 5 requests.
- on the 5th collision, the square should return a 418 and the request should be considered failed.

the server can both send and receive requests. it should take a webhook, assign the request to a river queue that ensures there can be a max of x squares originating from either side.

an example request body:
```json
{
    "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
    "created_at": "2024-09-07 04:00:00+10:00",
    "origin": "top",
    "colour": "#C8A2C8",
    "width": 5
}
```

### structure

these will be the main packages the app is composed of.

#### sender 

should be a standalone package that has a cmd you can run to send requests from the terminal.

#### server 

http server that sends and receives requests. 

#### model 

generated sqlc models.
