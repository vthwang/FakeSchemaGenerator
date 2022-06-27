# FakeSchemaGenerator
1. Clone this project.
2. Setup env file. Create `.env` in the document root.
   ```
   INDEX_URL=http://index.murmurations.dev/v2/nodes
   GITHUB_PAGE_URL=https://tingsyuanwang.github.io/FakeSchemaGenerator
   ```
3. `make run` to generate fsg executable file.
4. `make generate n=11000` - generate 11,000 files.
5. `make post n=11000` - post 11,000 to the server.