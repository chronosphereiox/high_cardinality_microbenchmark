# high_cardinality_benchmarks

This shows an exagerated case of where the total number of unique pod values 
and such can slow down any non-exact match label matcher at high cardinality
depending on whether a prefix search is applied and also how the speed of 
intersecting postings lists can also affect search speed when a lot of distinct
time series is matched each having a distinct label value matched by a query.

*Note:* As mentioned elsewhere, this is not meant to show real world 
performance, this is more meant to show an example of the effect of different
approaches on a micro level.
