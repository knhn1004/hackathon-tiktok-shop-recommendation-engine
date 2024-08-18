# Hackathon - TikTok Idea 2 Version

## Tech Stack

### Languages
- Go
- Python
- TypeScript
- JavaScript
- CSS

### Frameworks
- Next.js
- Kafka

### Tools
- Docker

### Infrastructure
- Cloudflare
- MySQL

## Project Description

Develop a content recommendation engine for TikTok Shop that suggests relevant products or content to users based on keywords and tags, displayed upon login.

### Architecture

**Backend**:  
- Golang server handles user sessions and content recommendations.
- Python scripts are used for processing and analyzing keywords to provide more accurate recommendations.

**Data Pipeline**:  
- Kafka handles real-time data streaming and communication between services, ensuring that user interactions and content updates are processed efficiently.

**Frontend**:  
- The user interface is built with TypeScript and styled with CSS, utilizing Next.js to create a seamless and responsive experience for users browsing recommended content.

**Database**:  
- MySQL stores user data, product details, and recommendation results, ensuring that personalized suggestions are both accurate and up-to-date.

**Communication**:  
- gRPC is implemented for efficient and low-latency communication between the Golang backend and the Python-based recommendation service.

**Infrastructure**:  
- Cloudflare is used to enhance the security and performance of the application, providing reliable content delivery and protection against potential threats.

## How It Works

### Scroll through Videos
Each card represents a TikTok video. Within each video, specific products are embedded, allowing users to discover new items directly from the content they're viewing.

### Find Products You Love
As users interact with product pages, our recommendation engine, powered by a language model (LLM), learns from their preferences and suggests products that align with their interests. This personalization enhances the shopping experience by curating content that resonates with the user's tastes.

### Seamless Integration
The recommendation engine seamlessly integrates into the TikTok Shop experience, providing users with relevant suggestions without interrupting their browsing. Whether they're scrolling through videos or exploring product pages, the engine dynamically adjusts to offer the most pertinent content.




## Development
Start
```
docker-compose -f docker-compose.dev.yml up --build --force-recreate
```

Logs
```
docker-compose -f docker-compose.dev.yml logs -f frontend
```

Down
```
docker-compose -f docker-compose.dev.yml down
```

db:seed (important to see items in API endpoint)
```
docker-compose -f docker-compose.dev.yml exec api-node make seed
```


![System Design]("/system_design.png")