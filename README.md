# syncmap

Implements IntSyncMap and StrSyncMap


```
imap := syncmap.NewIntSyncMap()
imap.Store(75, obj)
obj, ok := imap.Load(75)

smap := syncmap.NewStrSyncMap()
smap.Store("foo", obj)
obj, ok := smap.Load("foo")

```

