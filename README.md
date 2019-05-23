# gof-art
A small web service that take an image from somewhere and create a nice ascii art version of it.

# Quick test

```bash
curl -H "Content-Type: application/json" \                                  ✔  6927  18:09:57
--request POST \
--data '{"url":"https://2.bp.blogspot.com/-50t8QbXgxwI/WGgpaXNAYWI/AAAAAAAAEPE/SKJ-Bu12qpkrP7kklk1_QmWTehLoBRFcwCLcB/s1600/Gophers.jpg","width":250}' \
http://localhost:9999/ascii > test.txt
```