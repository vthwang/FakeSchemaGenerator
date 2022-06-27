# FakeSchemaGenerator
1. Clone this project.
2. Setup env file. Create `.env` in the document root.
   ```
   INDEX_URL=http://index.murmurations.dev/v2/nodes
   GITHUB_PAGE_URL=https://tingsyuanwang.github.io/FakeSchemaGenerator
   ```
3. `make run` to generate fsg executable file.
4. `make generate n=10010` - generate 10,010 files.
5. Update data to Github.
6. `make post n=10010` - post 10,010 to the server.
