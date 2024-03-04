# video-streaming-platform
The Video Streaming Platform project is a comprehensive media streaming service emulation, designed to replicate the functionalities of popular streaming platforms. It encompasses a range of content, spanning movies, TV series, and more. The core objective of this project lies in the implementation of search and recommendation mechanisms, leveraging advanced data structures and algorithms to enhance user experience.

## Features:

### Feature #1: Typo-Tolerant Search with Anagram Matching

Employs a search algorithm capable of accommodating minor typos by utilizing anagrams on user-provided input. This feature ensures that users receive relevant search results even in the presence of typographical errors.

```
                Welcome to Video Streaming Platform
1. Search for a show
2. Display Top shows
3. Display recently watched
4. Exit
Select an option: 1
Enter show's title: 
spedE   
speed found.
```


### Feature #2: Global Top-Rated Movies Aggregation

Utilizes a LinkedList data structure to aggregate movie rankings from geographic regions into a unified ranking system. By amalgamating these rankings, users gain access to a curated list of the highest-rated movies globally.

![Screenshot from 2024-03-04 08-21-49](https://github.com/fynecontry/video-streaming-platform/assets/27024731/32c2be4a-c326-48ee-8762-9761b7c0bcb6)

```
$ go run .
                Welcome to Video Streaming Platform
1. Search for a show
2. Display Top shows
3. Display recently watched
4. Exit
Select an option: 2
Tenet           Ratings: 1
King            Ratings: 11
Barbie          Ratings: 23
Code8           Ratings: 23
Inception               Ratings: 32
Speed           Ratings: 38
Maverick                Ratings: 44
Frozen          Ratings: 57
Jumanji         Ratings: 99
```


### Feature #3: Cache Management with LRU Replacement Strategy

Implements a caching mechanism augmented with a Least Recently Used (LRU) replacement strategy. This strategy ensures optimal utilization of system resources by replacing the least recently watched title in the cache with recently accessed titles, thereby enhancing the efficiency of content delivery.

![image](https://github.com/fynecontry/video-streaming-platform/assets/27024731/3877e445-f321-460c-b310-31e1b7cde096)

```
Welcome to Video Streaming Platform
1. Search for a show
2. Display Top shows
3. Display recently watched
4. Exit
Select an option: 3
The most recently watched titles are: (Show, rank)
(Barbie,23)

(Barbie,23)(King,11)

(Barbie,23)(King,11)(Inception,32)

(King,11)(Inception,32)(Code8,23)

(Inception,32)(Code8,23)(Tenet,1)

(Inception,32)(Code8,23)(Tenet,1)
```



This project highlights the utilization of advanced data structures and algorithms in the context of media streaming services.