package main

const (
	Script1 = `
		local value = redis.call("Get", KEYS[1])
		if( value - KEYS[2] >= 0 ) then
			local leftStock = redis.call("DecrBy" , KEYS[1],KEYS[2])
			return leftStock
		else
			return value - KEYS[2]
		end
	`
	//判断一个人的年龄字段是否存在，存在则加1，不存在则赋值为18
	Script2 = `
		local value = redis.call("HEXISTS",KEYS[1],KEYS[2])
		if (value == 1) then
			local rst = redis.call("HINCRBY",KEYS[1],KEYS[2],KEYS[3])
			return rst
		else
			local rst = redis.call("HSET",KEYS[1],KEYS[2],KEYS[4])
			return KEYS[4]
		end
	`
	Script3 = `
		function sleep(n)
			local t0 = os.clock()
			while os.clock - t0 <=n do end
		end
		local value = redis.call("HEXISTS",KEYS[1],KEYS[2])
		sleep(10000)
		if (value == 1) then
			local rst = redis.call("HINCRBY",KEYS[1],KEYS[2],KEYS[3])
			return rst
		else
			local rst = redis.call("HSET",KEYS[1],KEYS[2],KEYS[4])
			return KEYS[4]
		end
	`
)
