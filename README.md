# website

## Install Hugo

Version `v0.85.0` is used by our Netlify deployment container. Therefore, run:

```shell
CGO_ENABLED=1 go install -tags extended github.com/gohugoio/hugo@v0.85.0
```

Verify installation

```shell
hugo version
```

## New Blog Post

1. Copy one of the existing Markdown files in `content.en/blog/` and rename it to your liking. Make sure to prefix it with the rough date.
2. Remove all the content and only edit the frontmatter
   ```yml
   ---
   title: "Introducing Probelab"   # the fat title at the top
   date: 2024-01-11T12:33:50+01:00 # relevant for ordering
   slug: introducing-probelab      # this will become the path component 
   ---
   ```

The list of blog posts at `https://probelab.io/blog/` will automatically be updated to also include the new post. 
