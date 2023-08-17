# Backend Assignment (Intern) | FamPay

## Project Goal

To create an API for fetching the latest videos from YouTube, sorted in reverse chronological order of their publishing date-time, for the search query "F1". The API should provide paginated responses.

## Tech Stack Used

- Programming Language: Go (Golang)
- Web Framework: GoFiber
- Database: PostgreSQL

## Basic Requirements

- ☑️ Server continuously calls the YouTube API asynchronously with a 10-second interval to fetch and store the latest videos for the search query "F1". The video data, including fields like video title, description, publishing datetime, thumbnail URLs, and more, is stored in a database with proper indexes.
- ☑️ A GET API returns stored video data in a paginated response, sorted in descending order of published datetime.
- ☑️ The solution should be scalable and optimized for performance.

## Bonus Points

- ☑️ Add support for multiple API keys to automatically switch to the next available key when the quota is exhausted.
- ☑️ Create a dashboard to view stored videos with filter and sorting options (optional).

## Deployment Information
-The backend of this project is deployed on an AWS EC2 instance. The deployment ensures that the API is accessible and functional in a production environment.

-To enhance scalability, the Docker Compose configuration can be adjusted to increase the number of replicas in the web service. This helps in balancing the load and ensuring efficient handling of incoming requests.
## Deployment link
### Base Url: 
http://43.205.231.225/
### Postman Docs
https://documenter.getpostman.com/view/16178117/2s9Y5R37FY

## Instructions

### Running Instructions
- To build and run
```shell
docker-compose up -d --build
```
## Efficient Data Search with tsvector and GIN Index

In your project, the PostgreSQL database is optimized for effective text-based searches using `tsvector` and the `GIN` index.

- **tsvector Data Type**: The `tsvector` data type processes text, transforming words into normalized lexemes. It enables efficient searches by generating processed text representations.

- **GIN Index**: The `GIN` index is a specialized index for full-text searches. It's applied to the `tsvector` column, allowing PostgreSQL to swiftly retrieve matching documents.

In your implementation:

1. Video titles and descriptions are processed using `to_tsvector` and stored in the `search_weights` column.
2. User queries are converted into `tsquery` using `plainto_tsquery`.
3. Searches are conducted using the `@@` operator on the indexed `tsvector`.
4. The `search__weights_idx` GIN index optimizes search performance.

These techniques ensure rapid and precise search results based on title and description content.
