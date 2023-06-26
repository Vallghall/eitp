/**
 * Узел списка
 */
class node<T> {
  public value: T
  public next: node<T> | null

  constructor(value: T) {
    this.value = value
    this.next = null
  }
}

/**
 * Обертка/фиктивный узел списка
 */
export class List<T> extends Object {
  private _length: number
  private _head: node<T> | null

  constructor(...args: T[]) {
    super()
    this._length = args.length
    if (this._length == 0)
      return

    this._head = new node(args[0])
    let cur = this._head
    for (let i = 1; i < this._length; i++) {
      cur!.next = new node(args[i])
      cur = cur!.next
    }
  }

  public get length() {
    return this._length
  }

  public override toString(): string {
    let s = "["
    if (this._length > 0)
      s += this._head!.value

    let cur = this._head!.next
    while (cur != null) {
      s += "," + cur.value
      cur = cur.next
    }

    return s + "]"
  }

  /**
   * Добавляет элемент в голову списка
   */
  public prepend(value: T): List<T> {
    this._length++
    const n = new node(value)
    n.next = this._head
    this._head = n

    return this
  }

  /**
   * Добавляет элемент в конец списка
   */
  public append(value: T): List<T> {
    const n = new node(value)
    this._length++

    if (this._head == null) {
      this._head = n
      return this
    }

    let cur = this._head
    while (cur.next != null) {
      cur = cur.next
    }

    cur.next = n
    return this
  }

  /**
   * addAt добавляет элемент по данному индексу
   * @throws генерирует исключение, если индекс больше или равен длине массива
   */
  public addAt(idx: number, value: T): List<T> {
    if (idx < 0 || idx > this._length)
      throw new Error(`index out of range - {this.length}`)

    if (idx == 0) return this.prepend(value)
    if (idx == this.length) return this.append(value)

    let cur = this._head!.next
    for (let i = 1; i < idx - 1; i++)
      cur = cur!.next

    const n = new node(value)
    n.next = cur!.next
    cur!.next = n
    this._length++

    return this
  }

  /**
   * shift возвращает первый элемент списка и удаляет его
   */
  public shift(): T | null {
    if (this._head == null)
      return null

    const result = this._head!.value
    this._head = this._head!.next
    this._length--

    return result
  }
  
  /**
   * pop возвращает последний элемент списка и удаляет его
   */
  public pop(): T | null {
    if (this._head == null)
      return null

    this._length--
    
    if (this._head.next == null) {
      let result = this._head.value
      this._head = null
      return result
    }

    let prv: node<T> | null = this._head
    let cur = prv.next
    while (cur!.next != null){
      prv = cur
      cur = cur!.next
    }

    prv!.next = null
    return cur!.value
  }

  /**
   * delete возвращает элемент по данному индексу и удаляет его из списка
   * @throws генерирует исключение, если индекс больше или равен длине массива
   */
  public delete(idx: number): T | null {
    if (idx < 0 || idx >= this._length)
      throw new Error(`index out of range - {this.length}`)

    if (idx == 0) return this.shift()

    let prv = this._head
    let cur = prv!.next
    for (let i = 1; i < idx; i++){
      prv = cur
      cur = cur!.next
    }

    prv!.next = cur!.next
    return cur!.value
  }

  /**
   * head возвращает первый элемент списка, не удаляя его
   */
  public head = (): T | null => this._head!.value

  /**
   * tail возвращает хвост списка
   */
  public tail = (): List<T> => Object.setPrototypeOf({
    _length: this.length - 1,
    _head: this._head!.next
  }, List<T>)

  /**
   * get возвращает элемент по данному индексу
   * @throws генерирует исключение, если индекс больше или равен длине массива
   */
  public get(idx: number): T | null {
    if (idx < 0 || idx >= this._length)
      throw new Error(`index out of range - {this.length}`)
    
    let cur = this._head
    for (let i = 0; i < idx; i++)
      cur = cur!.next

    return cur!.value
  }

  /**
   * get возвращает элемент по данному индексу
   * @throws генерирует исключение, если индекс больше или равен длине массива
   */
  public set(idx: number, value: T): T | null {
    if (idx < 0 || idx >= this._length)
      throw new Error(`index out of range - {this.length}`)
    
    let cur = this._head
    for (let i = 0; i < idx; i++)
      cur = cur!.next

    return cur!.value = value
  }

  /**
   * Разворот списка
   */
  public reverse(): List<T> {
    if (this.length < 2)
      return this

    let prv: node<T> | null = null
    let cur = this._head
    let nxt = cur!.next
    while (cur!.next != null) {
      cur!.next = prv
      prv = cur
      cur = nxt
      nxt = nxt!.next
    }

    cur!.next = prv
    this._head = cur

    return this
  }
}
