## Exercise 5: Pub/Sub Event System

### Description
Build a publish-subscribe event system where publishers broadcast messages and subscribers receive them. The system should:
- Support multiple publishers
- Support multiple subscribers
- Allow subscribers to filter messages by topic
- Handle dynamic subscription/unsubscription
- Prevent blocking publishers

### Requirements
- Implement a `PubSub` system with:
  - `Subscribe(topic)` - returns a channel for that topic
  - `Publish(topic, message)` - sends to all subscribers of that topic
  - `Unsubscribe(topic)` - close and remove subscriber channel
- Support wildcard topics (e.g., "weather.*" matches "weather.nyc", "weather.london")
- Use `select` to avoid blocking when publishing
- Ensure subscribers don't block publishers
- Handle subscriber disconnection gracefully
- Implement backpressure handling (buffer limits)
- Create 3 publishers and 5 subscribers with different topic interests

### Example Topics
- "weather.nyc"
- "weather.london"
- "news.tech"
- "news.sports"
- "system.alert"

### Expected Output
```
Subscriber 1 subscribed to: weather.*
Subscriber 2 subscribed to: news.*
Subscriber 3 subscribed to: weather.nyc
Subscriber 4 subscribed to: system.*
Subscriber 5 subscribed to: news.tech

Publisher 1 published to weather.nyc: "Sunny, 72°F"
  -> Received by Subscriber 1
  -> Received by Subscriber 3

Publisher 2 published to news.tech: "New Go 1.22 released!"
  -> Received by Subscriber 2
  -> Received by Subscriber 5

Publisher 3 published to system.alert: "High CPU usage"
  -> Received by Subscriber 4

[Unsubscribing Subscriber 3 from weather.nyc]
Publisher 1 published to weather.nyc: "Cloudy, 65°F"
  -> Received by Subscriber 1
```

### Challenge
- Add message persistence (history for new subscribers)
- Implement at-least-once delivery guarantees
