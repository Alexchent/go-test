-- redis实现限流器
local key1 = KEYS[1]
local expire = ARGV[1]
local count = redis.call('incr', key1)
if (redis.call('ttl', key1) == -1) then
    redis.call('expire', key1, expire)
end

if tonumber(count) > tonumber(ARGV[2]) then
    return 0
else
    return 1
end