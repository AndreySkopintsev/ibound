module baseless

go 1.24.4

replace api => ./api

replace cache => ./cache

require api v0.0.0-00010101000000-000000000000

require cache v0.0.0-00010101000000-000000000000 // indirect

require github.com/gorilla/mux v1.8.1 // indirect
