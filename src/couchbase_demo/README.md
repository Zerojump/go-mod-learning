https://dzone.com/articles/getting-started-with-couchbase-using-docker-1
https://docs.couchbase.com/server/6.0/introduction/intro.html

docker run -d -p 8091:8091 couchbase

http://$IP:8091/

Do not increment or decrement counters if using XDCR

Note that since the append operation is done atomically, there is no need for a CAS check.
Users of the append and prepend operations should ensure that the resulting documents do not become too large. Couchbase has a hard document size limit of 20MB.
Using append and prepend on larger documents may cause performance degradation and memory fragmentation at the server level, as for each append operation the server must allocate memory for the new document size and then append the fragment to the new memory. The performance impact may be significant when document sizes reach beyond 100KB.
Finally, note that while append saves network traffic from the client to server (by only specifying the fragment to append), the entire document is replicated for each mutation. Five append operations on a single 10MB document will result in 50MB of traffic to each replica.
