## Installing Bookbank

#### This guide teaches you how to install Bookbank and get it running for the first time.

---
### Ways to install

- [VIA docker cli](#install-with-docker-cli)
- [VIA docker compose](#install-with-docker-compose)
---

#### Pre install:

- Create necessary files and folders:

```bash
mkdir -p bookbank bookbank/internal bookbank/internal/files
echo "{\"books\": []}" > bookbank/internal/books.json
echo "0" > bookbank/internal/counter
```

---

#### Install with Docker CLI

- Run:

```bash
docker run --publish 8080:8080 --restart unless-stopped \
-v $(pwd)/bookbank/internal/books.json:/app/helper/books.json \
-v $(pwd)/bookbank/internal/counter:/app/helper/counter \
-v $(pwd)/bookbank/internal/files:/app/store \
--name bookbank alwaysnur/bookbank
```

---

#### Install with Docker Compose

- Create a `compose.yaml` file and copy the contents of the
  [compose file provided](./examples/compose.yaml) into it.
- Then run `docker compose up`.

---
