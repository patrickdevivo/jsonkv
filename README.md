## jsonkv 🌲

Convert lists of KV pairs to JSON and vice versa.

Calling `ToKVs` with a JSON blob like this:
```js
{
  "key": {
    "at" : {
      "some": {
        "depth": true
      }
    }
  },
  "at_root": 123
}
```

Will return something like this:
```js
[
  {"key": "key/at/some/depth", "value": "true"},
  {"key": "at_root", "value": "123"},
]
```

And `ToJSON` will do the inverse. Note that JSON -> KVs turns all "leaves" of the JSON tree (numeric, boolean, string, null values) to strings.
