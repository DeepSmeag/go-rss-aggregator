# Go RSS Aggregator

Following tutorial at [link](https://www.youtube.com/watch?v=un6ZyFkqFKo&t=31155s)

Goal is to get more hands-on experience with Go while building this project. Especially handling a database connection.

As pe video chapters, it contains:

- Project description & setup
- Chi router usage
- Postgres usage
- Authentication w/ API keys
- Many to many relationships in db
- Aggregation worker
- Viewing the blog posts

## What I've learned

- lowercase/uppercase naming conventions for exporting functions to other packages applies to struct fields as well; the json.Marshal function could only serialize the fields if they were uppercase; similarly, if I'm trying to use/initialize a struct in another package, I only get the uppercase fields; this one's unintuitive
- and also learned that some people are in favor of commiting the vendor folder, since it's not that big; I don't understand the underlying reason enough to have an opinion on this yet

  - it has to do with ensuring the entire dev team works with the same dependencies, but I thought go.sum and go.mod files were enough for that
  - or maybe it's like a .lock (npm) file where instead of storing versions, we store the code itself so that there's no chance something changes in the library and it blocks us
  - the issue with this thinking is that we have package versions; aren't they supposed to ensure there's no unwanted update that potentially breaks our code?
  - I'm not committing the vendor folder since it's a small project; in bigger projects I'd think about it some more / follow the team's convention

- for the postgres part, I'm running off of docker with pgadmin for the GUI
- we're going with sqlc (link in video outdated) and goose (link in video ok); sqlc is a code generator that generates Go code from SQL queries; goose is a database migration tool; the video uses "go install" to download and build them

  - go install requires having $GOBIN in the $PATH; on macOS brew is recommended to install; manually adding the path creates the need to manually cleanup after uninstalling sqlc or goose so I'd not recommend that

- using SQLC & GOOSE
  - goose's role is to manage the database schema, so the structure of tables in the database; we manually write the SQL queries which create the tables or alter them, or undo changes (drop / alter)
  - sqlc's role then is to hook up to goose and the SQL code we write to generate Go code (type safe) automatically, so we can then use the generated Go code to interact with the database
  - // this is a more manual process compared to ORMs (I mainly used Drizzle in Node.js), but it the benefit is we have way more control over the SQL queries and how our database changes;
  - // especially for migrations when a resource schema needs changing, ORMs are a bit harder to understand in what they're going to do to existing data; with goose we can choose that ourselves
- a good idea is to have a package handling DTOs, or rather something that converts the database data to structs we want to use in our code / push to the end user; simplest example: don't give the user their internal id;
- important note: this project doesn't necessarily follow a scalable structure; the handlers are in the main package after all (and the middleware is structured in a way that forces it to be in the same package as the handlers)
