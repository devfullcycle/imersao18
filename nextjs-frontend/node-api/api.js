const express = require("express");
const bodyParser = require("body-parser");
const { events, spots } = require("./data");

const app = express();

app.use(bodyParser.json());

app.get("/events", async (req, res) => {
  return res.json({
    events,
  });
});

app.get("/events/:eventId", async (req, res) => {
  const event = events.find(
    (event) => event.id === parseInt(req.params.eventId)
  );
  if (!event) {
    return res.status(404).json({
      message: "Event not found",
    });
  }
  return res.json(event);
});

app.get("/events/:eventId/spots", async (req, res) => {
  const eventId = parseInt(req.params.eventId);
  const event = events.find((event) => event.id === eventId);
  if (!event) {
    return res.status(404).json({
      message: "Event not found",
    });
  }
  const eventSpots = spots.filter((spot) => spot.event_id === eventId);
  return res.json({
    event,
    spots: eventSpots,
  });
});

app.post("/checkout", async (req, res) => {
  const { event_id, card_hash, ticket_kind, spots: spotsName, email } = req.body;
  console.log(req.body);
  const event = events.find((event) => event.id == event_id);
  if (!event) {
    return res.status(404).json({
      message: "Event not found",
    });
  }
  if (!card_hash) {
    return res.status(400).json({
      message: "Card hash is required",
    });
  }
  if (!ticket_kind) {
    return res.status(400).json({
      message: "Ticket kind is required",
    });
  }
  if (!spotsName) {
    return res.status(400).json({
      message: "Spots are required",
    });
  }
  if (!email) {
    return res.status(400).json({
      message: "Email is required",
    });
  }
  const eventSpots = spots.filter((spot) => spot.event_id == event_id);
  for (const spotName of spotsName) {
    const spotIndex = eventSpots.findIndex((s) => s.name === spotName);
    
    if (spotIndex === -1) {
      return res.status(400).json({
        message: `Spot ${spot.id} not found`,
      });
    }
    if (eventSpots[spotIndex].status !== "available") {
      return res.status(400).json({
        message: `Spot ${eventSpots[spotIndex].id} is not available`,
      });
    }
    eventSpots[spotIndex].status = "sold";
  }

  return res.json({
    message: "Spots reserved successfully",
  });
});

const port = 8080;
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
