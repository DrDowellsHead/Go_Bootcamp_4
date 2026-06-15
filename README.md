# Project 03 — Go_Bootcamp

**Overview**: In this project, you will learn how to build a web application in Go using the net/http package.  
💡 [Click here](https://new.oprosso.net/p/4cb31ec3f47a4596bc758ea1861fb624) to share your feedback on this project. It's anonymous and will help our team improve the learning experience. We recommend completing the survey right after finishing the project.

## Contents

  - [Chapter I](#chapter-i)
    - [Instructions](#instructions)
  - [Chapter II](#chapter-ii)
    - [General Information](#general-information)
  - [Chapter III](#chapter-iii)
  - [Project: Tic-Tac-Toe](#project-tic-tac-toe)
    - [Task 0. Project Initialization](#task-0-project-initialization)
    - [Task 1. Project Structure](#task-1-project-structure)
    - [Task 2. Implement the Domain Layer](#task-2-implement-the-domain-layer)
    - [Task 3. Implement the Datasource Layer](#task-3-implement-the-datasource-layer)
    - [Task 4. Implement the Web Layer](#task-4-implement-the-web-layer)
    - [Task 5. Implement the DI Layer](#task-5-implement-the-di-layer)


## Chapter I

### Instructions

How to learn at “School 21”:

- Here, you’ll find a unique learning experience with a lot of freedom. You’re given a task and left to find your own way to solve it, using whatever resources work best for you — whether that’s the Internet or AI tools like GigaChat. Just be mindful of information quality: verify, think critically, analyze, and compare.
- Peer-to-peer (P2P) learning is the exchange of knowledge and experience with peers, where everyone acts as both mentor and student. This approach allows you to gain a deeper understanding of the material by learning from one another.
- Feel free to ask for help: around you are peers who are also navigating this path for the first time. Share your own experience and ideas with others.  Join Rocket.Chat to stay updated with the latest community announcements. 
- Your learning is meaningless if you just copy someone else’s solutions. When receiving help from others, always make sure you fully understand the “why”, “how”, and “purpose” behind the solution. Don’t be afraid to make mistakes. 
- Does the task seem impossible? Take a break, get some fresh air and clear your mind — this has helped many people. Maybe after that, the solution will come to you naturally.
- The learning process is just as important as the result. It’s not just about completing the task — it’s about understanding HOW to solve it. 

How to work with the project:

- Before starting, clone the project from GitLab into a repository with the same name.
- All files should be created inside the _src/_ folder of the cloned repository.
- After cloning the project, create a _develop_ branch and do all your development there. Then, push the _develop_ branch to GitLab.
- Your directory should not contain any files other than those specified in the assignments.

## Chapter II

### General Information

A web application is a client-server application where the client interacts with the web server using a browser. The application logic is distributed between the server and the client; data is stored primarily on the server, and communication happens over a network.

The net/http package is a powerful standard Go library for building web applications.

Advantages of net/http:

- **Ease of use**: net/http has a simple and intuitive syntax, making it ideal for beginners or developers who prefer clean and readable code.
- **Maintainability**: Any network-related task can be solved either using net/http itself or with community-built packages designed for it.

These advantages make net/http an appealing choice for both beginners and advanced developers looking to build fast, scalable, and maintainable web applications in Go.

**Topics to study**:

- Web application
- net/http
- API
- Minimax algorithm
- MVC

## Chapter III

## Project: Tic-Tac-Toe

### Task 0. Project Initialization
You need to create a project that will be used for all subsequent tasks.

### Task 1. Project Structure

- Each layer should be its own module.
- The project structure must follow the **Standard Go Project Layout**.
- Separate the **API contracts layer** for client interaction.
- Separate the **application layer** for business logic implementation.
- Separate the **infrastructure layer** for handling data operations (e.g., with a database).
- Separate the **DI (dependency injection) layer**, where configurations for dependency injection are described (e.g., using the uber/fx library).

### Task 2. Implement the Domain Layer

- Define a **game board model** as an integer matrix.
- Define a **current game model**, which should include a UUID and the game board.
- Define a **service interface** with the following methods:
  - A method to compute the next move of the current game using the **Minimax algorithm**.
  - A method to **validate** the game board of the current game (check that previous moves haven't been altered).
  - A method to **check for game completion**.
- Models, interfaces, and implementations must be placed in **separate files**.

### Task 3. Implement the Datasource Layer

- Implement a **storage structure** to keep track of ongoing games.
- Use **thread-safe collections** (e.g., sync.Map?) for storing data.
- Define models for the current game’s game board.
- Implement **mappers** between domain and datasource layers (domain <-> datasource).
- Implement a **repository** to interact with the storage structure. It must provide the following methods:
  - A method to **save the current game**.
  - A method to **retrieve the current game**.
- Create a structure that implements the **service interface** and accepts a **repository interface** as a parameter to work with the storage structure.
- Models, interfaces, and implementations must be placed in **separate files**.

### Task 4. Implement the Web Layer

- Define **models** for the current game’s game board.
- Implement **mappers** between the domain and web layers (domain <-> web).
- Implement a **handler** using net/http, with a method:
  - POST /game/{current_game_UUID} — sends the current game with the user’s updated game board and returns the current game with the computer’s updated game board.
- If an **invalid game** with an incorrect updated board is sent, an **error with a description** must be returned.
- The application must support **multiple games simultaneously**.
- Models, interfaces, and implementations must be placed in **separate files**.

### Task 5. Implement the DI Layer

- Use the uber/fx library to define a **dependency injection graph**.
- The graph must include at least the following components:
  - The **storage structure**, registered as a **singleton**;
  - The **repository** for working with the storage structure;
  - The **service** for working with the repository.