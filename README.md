# socialnetwork

(Chat made...)

Frontend:
1. Setup:

    Choose a JavaScript framework (React, Vue.js, Svelte, or Mithril).
    Set up the project structure.

2. User Authentication:

    Create registration and login forms.
    Implement sessions and cookies for keeping users logged in.
    Add a logout option.

3. Profile Page:

    Design and create a user profile page.
    Display user information (except password).
    Display user activity (posts).
    Implement UI for followers and following.
    Toggle for making profile public or private.

4. Posts:

    Implement creating, editing, and deleting posts.
    Allow image or GIF attachments in posts.
    Implement comments on posts.
    Privacy options for posts (public, private, custom).

5. Groups:

    Implement group creation, joining, and invitations.
    Create group browsing section.
    Design group page (posts, comments, members).
    Implement event creation within groups, including RSVP options.

6. Followers:

    Implement UI for sending and receiving follow requests.
    Display followers and following on profile.

7. Chats:

    Implement UI for private messaging between users.
    Include emoji support in chat.
    Implement group chat for groups.

8. Notifications:

    Implement UI for displaying notifications.
    Separate UI elements for chat messages and other notifications.

9. Responsiveness and Performance:

    Ensure the website is responsive on various devices.
    Optimize performance.

Backend:
1. Setup:

    Choose a web server (e.g. Caddy) or create your own.
    Set up the project structure.
    Set up SQLite database and create an Entity Relationship Diagram.

2. User Authentication:

    Implement user registration and login logic.
    Implement sessions and cookies for authentication.

3. Profile:

    Implement endpoints to retrieve and update user profiles.
    Implement logic for public and private profiles.

4. Posts:

    Implement endpoints for creating, retrieving, updating, and deleting posts.
    Implement logic for post visibility based on privacy settings.
    Implement comments logic.

5. Groups:

    Implement endpoints for creating, joining, and managing groups.
    Implement group invitation logic.
    Implement event creation within groups, including RSVP options.

6. Followers:

    Implement endpoints for sending, accepting, and removing follow requests.
    Implement logic for followers and following.

7. Chats:

    Implement WebSocket endpoints for real-time chat.
    Store chat history in the database.

8. Notifications:

    Implement endpoints for retrieving notifications.
    Implement logic for generating notifications based on user actions.

9. Images Handling:

    Implement middleware for handling image uploads.
    Support JPEG, PNG, and GIF types.

10. Migrations:

    Implement migrations to create database tables.
    Use a package like golang-migrate for managing migrations.

11. Docker:

    Create Docker images for backend and frontend.
    Set up Docker Compose for local development.

12. Testing and Debugging:

    Write unit and integration tests for backend logic.
    Debug and resolve issues.
