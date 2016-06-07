
micro query go.micro.srv.gpsloc GpsLocation.Get '{"Id": "id123"}'

micro query go.micro.srv.gpsloc GpsLocation.Search '{"center": {"latitude": 51.5
165, "longitude": 0.1246, "timestamp": 145443}, "radius": 2, "type": "Hotel", "numEntities": 1}'

micro query go.micro.srv.gpsloc GpsLocation.Post '{"entity": {"location": {"lati
tude": 80.5165, "longitude": 0.2246, "timestamp": 11145}}, "name": "Changi", "type": "Airport"}'
{}
