-------------------------------------------study_room----------------------------------------
curl -i -X POST \
-H "Content-Type:application/json" \
-d \
'{
"name":"加油鸭自习室",
"intro":"介绍"
}' \
'http://127.0.0.1:8080/study_room/add'

curl -i -X POST \
-H "Content-Type:application/json" \
-d \
'{
"id": 1,
"name":"加油鸭自习室1",
"intro":"介绍1"
}' \
'http://127.0.0.1:8080/study_room/update'

curl -i -X POST \
-H "Content-Type:application/json" \
-d \
'{
"id": 1
}' \
'http://127.0.0.1:8080/study_room/del'
--------------------------------------------study_room------------------------------------





-------------------------------------------seat_type----------------------------------------
curl -i -X POST \
-H "Content-Type:application/json" \
-d \
'{
"sid":1,
"intro": "介绍",
"name": "舒适房间",
"price_intro": "{\"按日缴费\": 9,\"按日缴费\": 59,\"按日缴费\": 269}"
}' \
'http://127.0.0.1:8080/seat_type/add'


curl -i -X POST \
-H "Content-Type:application/json" \
-d \
'{
"id":2,
"sid":1,
"intro": "介绍",
"name": "舒适房间",
"price_intro": "{\"按日缴费123\": 9,\"按日缴费\": 59,\"按日缴费\": 269}"
}' \
'http://127.0.0.1:8080/seat_type/update'

curl -i -X POST \
-H "Content-Type:application/json" \
-d \
'{
"id": 1
}' \
'http://127.0.0.1:8080/seat_type/del'
-------------------------------------------seat_type----------------------------------------







-------------------------------------------seatinfo----------------------------------------
curl -i -X POST \
-H "Content-Type:application/json" \
-d \
'{
"sid": 1,
"seat_type_id": 1
}' \
'http://127.0.0.1:8080/seatinfo/add'

curl -i -X POST \
-H "Content-Type:application/json" \
-d \
'{
"id":2,
"seat_type_id": 2
}' \
'http://127.0.0.1:8080/seatinfo/update'

curl -i -X POST \
-H "Content-Type:application/json" \
-d \
'{
"id":1
}' \
'http://127.0.0.1:8080/seatinfo/del'
-------------------------------------------seatinfo----------------------------------------

-------------------------------------------buy_record----------------------------------------

-------------------------------------------buy_record----------------------------------------