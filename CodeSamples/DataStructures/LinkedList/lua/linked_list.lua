-- узел связанного списка
local node = {}
node.__index = node

function node:new(v)
    return setmetatable({value = v}, self)
end

-- обертка/фиктивный узел списка
List = {}
List.__index = List

function List:new()
    return setmetatable({head = nil, length = 0}, self)
end

-- метаметод для печати объекта списка
function List:__tostring()
    local str = "["

    if self.head ~= nil then
        str = str .. self.head.value
        local cur = self.head.next
        while cur do
            str = str .. "," .. cur.value
            cur = cur.next
        end
    end

    return str .. "]"
end

-- вставка в начало списка
function List:prepend(v)
    local n = node:new(v)
    n.next = self.head
    self.head = n
    self.length = self.length + 1

    return self
end

-- вставка в конец списка
function List:append(v)
    local n = node:new(v)
    self.length = self.length + 1

    if self.head == nil then
        self.head = n
        return
    end
    
    local cur = self.head
    
    while cur.next ~= nil do
        cur = cur.next
    end

    cur.next = n
    return self
end

-- метаметод для получения длины списка через
-- унарный оператор #
function List:__len()
    return self.length
end

-- добавление элемента в указанную позицию в списке
function List:addAt(idx, val)
    if idx < 0 or idx > #self then
        error("index" .. idx .. "is out of range")
    end

    if idx == 0 then
        return self:prepend(val)
    end

    if idx == #self then
        return self:append(val)
    end

    local cur = self.head
    for i = 1,idx - 1,1 do
        cur = cur.next
    end

    local n = node:new(val)
    n.next = cur.next
    cur.next = n

    return self
end

-- удаление и возврат элемента из начала списка
function List:shift()
    if self.head == nil then
        return nil
    end

    local val = self.head.value
    self.head = self.head.next
    self.length = self.length - 1
    
    return val
end

-- удаление и возврат элемента из конца списка
function List:pop()
    if self.head == nil then
        return nil
    end

    self.length = self.length - 1

    if self.head.next == nil then
        local val = self.head.value
        self.head = nil
        return val
    end

    local prv = self.head
    local cur = prv.next
    repeat
        prv = cur
        cur = cur.next
    until cur.next == nil

    prv.next = nil
    return cur.value
end

-- удаление и возврат элемента по индексу
function List:delete(idx)
    if idx < 0 or idx >= #self then
        error("index" .. idx .. "is out of range")
    end

    if idx == 0 then
        return self:shift()
    end

    self.length = self.length - 1

    local prv = self.head
    local cur = prv.next
    for i = 1,idx - 1,1 do
        prv = cur
        cur = cur.next
    end

    prv.next = cur.next
    return cur.value
end


-- возврат значения из головы списка
function List:head()
    return self.head.value
end

-- возврат хвоста списка
function List:tail()
    if self.head == nil then
        return nil
    end

    local list = List:new()
    list.head = self.head.next
    list.length = #self - 1

    return list
end

-- получение значения по индексу
function List:get(idx)
    if idx < 0 or idx >= #self then
        error("index " .. idx .. " is out of range")
    end

    local cur = self.head
    local i = 0
    while i < idx do
        i = i + 1
        cur = cur.next
    end

    return cur.value
end

-- установка значения по индексу
function List:set(idx, val)
    if idx < 0 or idx >= #self then
        error("index " .. idx .. " is out of range")
    end

    local cur = self.head
    local i = 0
    while i < idx do
        i = i + 1
        cur = cur.next
    end

    cur.value = val
end


-- разворот списка
function List:reverse()
    if #self < 2 then
        return
    end

    local cur = self.head
    local prv = nil
    local tmp = cur.next
    repeat
        cur.next = prv
        cur = tmp
        tmp = tmp.next
    until cur.next == nil
end

return List
