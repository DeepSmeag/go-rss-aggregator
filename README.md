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
