<h1 align="center">Bookbank</h1>
<h4 align="center">The audiobook/podcast media server solution.</h4>

![Showcase of bookbank](./assets/image.png)

### 🎧 Features

- Easy to use interface
- Free forever

### Install

With Docker CLI:

Create necessary files and folders

```bash
mkdir -p bookbank bookbank/internal bookbank/internal/files
echo "{\"books\": []}" > bookbank/internal/books.json
echo "0" > bookbank/internal/counter
```

```bash
docker run --publish 8080:8080 \
-v $(pwd)/bookbank/internal/books.json:/app/helper/books.json \
-v $(pwd)/bookbank/internal/counter:/app/helper/counter \
-v $(pwd)/bookbank/internal/files:/app/store \
alwaysnur/bookbank
```
