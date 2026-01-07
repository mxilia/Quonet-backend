[TYPESCRIPT_BADGE]: https://img.shields.io/badge/typescript-D4FAFF?style=for-the-badge&logo=typescript
[AWS_BADGE]:https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white
[GO_BADGE]:https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[REDIS_BADGE]:https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white
[POSTGRESQL_BADGE]:https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white
[SUPABASE_BADGE]:https://img.shields.io/badge/Supabase-3ECF8E?style=for-the-badge&logo=supabase&logoColor=white
[TAILWINDCSS_BADGE]:https://img.shields.io/badge/TailwindCSS-38BDF8?style=for-the-badge&logo=tailwindcss&logoColor=white
[NEXT_BADGE]:https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white

<h1 align="center" style="font-weight: bold;">Quonet</h1>
<p align="center">
 A forum web application that allows users to create, delete, and like or dislike posts. Users can also update their profile information, including their username and bio. The platform supports content moderation through an admin dashboard.
</p>

![Alt text](.github/assets/quonet_page.png)

<h3 align="center">
    <a href="https://www.quonet.dev"> www.quonet.dev</a>
</h3>

![typescript][TYPESCRIPT_BADGE]
![Go][GO_BADGE]
![AWS][AWS_BADGE]
![Redis][REDIS_BADGE]
![PostgreSQL][POSTGRESQL_BADGE]
![Supabase][SUPABASE_BADGE]
![TailwindCSS][TAILWINDCSS_BADGE]
![Next.js][NEXT_BADGE]

<h2 id="started">📁 Repositories</h2>

 - [Frontend](https://github.com/mxilia/Quonet-frontend)
 - [Backend](https://github.com/mxilia/Quonet-backend)

<h2 id="summary">📄 Summary</h2>

The frontend is built with Next.js following [Bulletproof Architecture](https://github.com/alan2207/bulletproof-react), leveraging server-side rendering for fast initial page loads, TanStack for client-side caching, and Zustand for centralized state management.

Backend is built with Go using Clean Architecture and implements a RESTful API with Fiber v2. PostgreSQL is used for data persistence via GORM, Redis handles rate limiting, and images are stored using Supabase.

<h2 id="tech">💻 Tech Stack</h2>

 - __Frontend:__ Next.js, ShadCN, Tailwind CSS, Zod, TanStack, Zustand
 - __Backend:__ Go, Fiber, GORM, PostgreSQL, Redis
 - __Service:__ AWS, Supabase

<h2 id="started">🚀 Getting started</h2>
<h3 id="prerequisites"> Prerequisites </h3>

 - Go 1.25.2
 - Docker
 - Docker Compose

<h3 id="setup"> Setting up </h3>

Run this to clone project:
```bash
git clone https://github.com/mxilia/Quonet-backend.git
```
After that make sure your current directory is at the root of this project.
```bash
cd Quonet-backend
```

Run this to download Go package:
```bash
go mod download
```

<h3 id="env">Environment Variables</h2>

Make sure you use .env.dev when development and .env.production in production. Here is the variable list example:
```yaml
ENV={dev or production}
DOMAIN=localhost

DB_HOST=localhost
DB_PORT=5432
DB_NAME=db_name
DB_USER=db_user
DB_PASSWORD=db_password

OWNER_EMAIL=xyz@email.com
OWNER_HANDLER=owner_handler

JWT_SECRET=nono22

FRONTEND_URL=https://localhost:3000
FRONTEND_OAUTH_REDIRECT_URL=http://localhost:3000

GOOGLE_CLIENT_ID=4567.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=4567
GOOGLE_OAUTH_REDIRECT_URL=http://localhost:8000/auth/google/callback

SUPABASE_URL=
SUPABASE_KEY=
BUCKET_NAME=

REDIS_ADDR=http://localhost:0000
REDIS_PASSWORD=
```
Write your own values for corresponding variables.

<h3 id="usage"> Usage </h3>

To run this project, you need to first open Docker Desktop then after that run this in terminal:
```bash
docker compose up --build
```

Then after that run:
```bash
go run cmd/app/main.go
```
And there you go, the backend is up and running.

<h3 id="note"> Note </h3>

📁 `/internal/app/server.go`
```
6 | db, storage, _, rateLimiter, cfg, err := setupDependencies(CURRENT_ENV)
```
Make sure to replace CURRENT_ENV according to your environment:
- Set `CURRENT_ENV = "production"` for a production environment.
- Set `CURRENT_ENV = "dev"` for a development environment.

<h2 id="structure">🧱 Project Structure</h2>

```
.
├── .github
│   └── workflows
│       └── cicd.yml
├── .vscode
│   └── settings.json
├── cmd
│   └── app
│       └── main.go
├── internal
│   ├── app
│   │   ├── app.go
│   │   └── server.go
│   ├── announcement
│   │   ├── dto
│   │   │   ├── mapper.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── handler
│   │   │   └── rest
│   │   │       └── handler.go
│   │   ├── repository
│   │   │   ├── announcement_repository.go
│   │   │   └── gorm_announcement_repository.go
│   │   └── usecase
│   │       ├── interface.go
│   │       └── usecase.go
│   ├── comment
│   │   ├── dto
│   │   │   ├── mapper.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── handler
│   │   │   └── rest
│   │   │       └── handler.go
│   │   ├── repository
│   │   │   ├── comment_repository.go
│   │   │   └── gorm_comment_repository.go
│   │   └── usecase
│   │       ├── interface.go
│   │       └── usecase.go
│   ├── entities
│   │   ├── announcement.go
│   │   ├── comment.go
│   │   ├── like.go
│   │   ├── post.go
│   │   ├── session.go
│   │   ├── thread.go
│   │   └── user.go
│   ├── like
│   │   ├── dto
│   │   │   ├── mapper.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── handler
│   │   │   └── rest
│   │   │       └── handler.go
│   │   ├── repository
│   │   │   ├── like_repository.go
│   │   │   └── gorm_like_repository.go
│   │   └── usecase
│   │       ├── interface.go
│   │       └── usecase.go
│   ├── post
│   │   ├── dto
│   │   │   ├── mapper.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── handler
│   │   │   └── rest
│   │   │       └── handler.go
│   │   ├── repository
│   │   │   ├── post_repository.go
│   │   │   └── gorm_post_repository.go
│   │   └── usecase
│   │       ├── interface.go
│   │       └── usecase.go
│   ├── session
│   │   ├── dto
│   │   │   ├── mapper.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── handler
│   │   │   └── rest
│   │   │       └── handler.go
│   │   ├── repository
│   │   │   ├── session_repostiory.go
│   │   │   └── gorm_session_repository.go
│   │   └── usecase
│   │       ├── interface.go
│   │       └── usecase.go
│   ├── thread
│   │   ├── dto
│   │   │   ├── mapper.go
│   │   │   ├── request.go
│   │   │   └── response.go
│   │   ├── handler
│   │   │   └── rest
│   │   │       └── handler.go
│   │   ├── repository
│   │   │   ├── thread_repository.go
│   │   │   └── gorm_thread_repository.go
│   │   └── usecase
│   │       ├── interface.go
│   │       └── usecase.go
│   ├── transaction
│   │   ├── interface.go
│   │   └── gorm_tx_manager.go
│   └── user
│       ├── dto
│       │   ├── mapper.go
│       │   ├── request.go
│       │   └── response.go
│       ├── handler
│       │   └── rest
│       │       └── handler.go
│       ├── repository
│       │   ├── user_repository.go
│       │   └── gorm_user_repository.go
│       └── usecase
│           ├── interface.go
│           └── usecase.go
├── pkg
│   ├── apperror
│   │   └── apperror.go
│   ├── config
│   │   └── config.go
│   ├── database
│   │   ├── db.go
│   │   └── storage.go
│   ├── middleware
│   │   ├── fiber.go
│   │   ├── jwt.go
│   │   ├── require_admin.go
│   │   └── rate_limit
│   │       ├── key.go
│   │       ├── policy.go
│   │       └── rate_limit.go
│   ├── redisclient
│   │   ├── client.go
│   │   └── kv.go
│   ├── responses
│   │   ├── error_response.go
│   │   └── message_response.go
│   ├── routes
│   │   ├── not_found_route.go
│   │   ├── private_routes.go
│   │   └── public_routes.go
│   └── token
│       ├── claims.go
│       └── jwt_maker.go
├── utils
│   └── format
│        └── format.go
├── .dockerignore
├── .env.dev
├── .env.example
├── .env.production
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

<h2 id="routes">📍 API Endpoints</h2>
<h3 id="api-v2"> /api/v2 </h3>
<table>
    <tr>
        <td align="center"><b>Route</b></th>
        <td align="center"><b>Auth Requirement</b></th>
        <td align="center"><b>Params</b></th>
        <td align="center"><b>Description</b></th>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /me</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Returns current user's info.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /auth/google/login</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Redirects user to Google OAuth Page.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /auth/google/callback</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Handles Google OAuth callback.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>POST /auth/logout</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Delete current user's session. 
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /threads</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            title, page
        </td>
        <td align="center">
            Returns threads that match params.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /threads/:id</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            id
        </td>
        <td align="center">
            Returns a thread with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>POST /threads</kbd>
        </td>
        <td align="center">
            Admin
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Creates a new thread.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>DELETE /threads/:id</kbd>
        </td>
        <td align="center">
            Admin
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Delete a thread with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>PATCH /sessions/:email</kbd>
        </td>
        <td align="center">
            Admin
        </td>
        <td align="center">
            email
        </td>
        <td align="center">
            Patch a session with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /users</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Returns all users.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /users/:id</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            id
        </td>
        <td align="center">
            Returns a user with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /users/handler/:handler</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            handler
        </td>
        <td align="center">
            Returns a user with specified handler.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /users/email/:email</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            email
        </td>
        <td align="center">
            Returns a user with specified email.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>PATCH /users/:id</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Update user with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>DELETE /users/:id</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Delete user with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /posts</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            page, author_id, thread_id, title
        </td>
        <td align="center">
            Returns posts that match params.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /posts/:id</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            id
        </td>
        <td align="center">
            Returns a post with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /posts/top/like</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            limit
        </td>
        <td align="center">
            Returns a limited number of posts sorted by like count in descending order.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /posts/private</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            page, author_id, thread_id, title
        </td>
        <td align="center">
            Returns private posts that match params.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /posts/private/:id</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            id
        </td>
        <td align="center">
            Returns a private post with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>POST /posts</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Creates a new post.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>PATCH /posts/:id</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            id
        </td>
        <td align="center">
            Update a post with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>DELETE /posts/:id</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            id
        </td>
        <td align="center">
            Delete a post with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /likes</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            page, parent_type, owner_id, parent_id
        </td>
        <td align="center">
            Returns likes that match params.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /likes/:id</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            id
        </td>
        <td align="center">
            Returns a like with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /likes/attributes/count</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            parent_type, parent_id, owner_id
        </td>
        <td align="center">
            Return a list of number representing entities' like count that match params.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /likes/attributes/state</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            parent_type, parent_id
        </td>
        <td align="center">
            Return current user like state for the given parent.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>POST /likes</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Create a new like.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /comments</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            root_id, parent_id, owner_id
        </td>
        <td align="center">
            Returns comments that match params.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /comments/:id</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            id
        </td>
        <td align="center">
            Returns a comment with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>POST /comments</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Create a new comment.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>PATCH /comments/:id</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Update a comment with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>DELETE /comments/:id</kbd>
        </td>
        <td align="center">
            Authenticated
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Delete a comment with specified id.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>GET /announcements</kbd>
        </td>
        <td align="center">
            Public
        </td>
        <td align="center">
            page
        </td>
        <td align="center">
            Returns announcements that match params.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>POST /announcements</kbd>
        </td>
        <td align="center">
            Admin
        </td>
        <td align="center">
            -
        </td>
        <td align="center">
            Create a new announcement.
        </td>
    </tr>
    <tr>
        <td align="center">
            <kbd>DELETE /announcements</kbd>
        </td>
        <td align="center">
            Admin
        </td>
        <td align="center">
            id
        </td>
        <td align="center">
            Delete an announcement with specified id.
        </td>
    </tr>
</table>