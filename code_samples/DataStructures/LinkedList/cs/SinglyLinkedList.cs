using System.Text;


// Обертка, выступающая в виде фиктивного узла
public class List<T> where T : IComparable<T>
{
    // Узел
    internal class Node 
    {
        public T? value; // Хранимое значение
        public Node? next; // Ссылка на следующий элемент цикла

        public Node(T? val)
        {
            value = val;
            next = null;
        }
    }
    internal Node? head; // указатель на первый элемент списка
    internal int length; // длина списка

    public int Length // свойство для доступа к значению длины
    {
        get => this.length;
    }

    public List()
    {
        this.head = null;
        this.length = 0;
    }

    // конструктор для составления списка из множества значений
    public List(params T?[] vals) : this()
    {
        if (vals is null || vals.Length == 0)
            return;

        this.head = new Node(vals[0]);
        Node? cur = this.head;

        for (var i = 1; i < vals.Length; i++) 
        {
            cur.next = new Node(vals[i]);
            cur = cur.next;
        }

        this.length = vals.Length;
    }

    // Добавление элемента в конец списка - медленно
    public List<T> Append(T? value)
    {
        this.length++;
        Node n = new(value);
        if (this.head is null)
        {
            this.head = n;
            return this;
        }

        Node current = this.head;
        while (current.next is not null)
            current = current.next;

        current.next = n;
        return this;
    }

    // Добавление элемента в начало списка - быстро
    public List<T> Prepend(T? value)
    {
        Node n = new(value);
        n.next = this.head;
        this.head = n;
        this.length++;

        return this;
    }

    // Добавление элемента в середину списка - зависит от позиции, но тоже медленно
    public List<T> AddAt(int idx, T? value)
    {
        if (idx < 0 || idx > this.Length)
                throw new System.IndexOutOfRangeException();

        if (idx == 0) return this.Prepend(value);

        if (idx == this.Length) return this.Append(value);

        Node cur = this.head;
        int j = 0;
        while (j < idx - 1)
        {
            cur = cur.next;
            j++;
        }

        Node nxt = new(value){next = cur.next};
        cur.next = nxt;
        this.length++;

        return this;
    }

    // Удаление элемента с начала списка
    public T? Shift()
    {
        if (this.head is null)
            return default(T);

        T result = this.head.value;
        this.head = this.head.next;
        this.length--;

        return result;
    }

    // Удаление элемента с конца списка
    public T? Pop()
    {
        if (this.head is null)
            return default(T);

        this.length--;
		
		if (this.head.next is null)
		{
			T result = this.head.value;
			this.head = null;
			return result;
		}

        Node prev = this.head;
		Node cur = prev.next;
        while (cur.next is not null) 
        {
            prev = cur;
            cur = cur.next;
        }

        prev.next = null;
        return cur.value;
    }

    public T? Delete(int idx)
    {
        if (idx < 0 || idx >= this.Length)
                throw new System.IndexOutOfRangeException();

        if (idx == 0) return this.Shift();

        Node prev = this.head;
		Node cur = prev.next;
        int j = 1;
        while (j < idx)
        {
            prev = cur;
            cur = cur.next;
            j++;
        }

        prev.next = cur.next;
        this.length--;

        return cur.value;
    }

    public T this[int idx]
    {
        get 
        {
            if (idx < 0 || idx >= this.length)
                throw new System.IndexOutOfRangeException();

            Node cur = this.head;
            int j = 0;
            while (j != idx)
            {
                cur = cur.next;
                j++;
            }

            return cur.value;
        }

        set
        {
            if (idx < 0 || idx >= this.length)
                throw new System.IndexOutOfRangeException();

            Node cur = this.head;
            int j = 0;
            while (j != idx)
            {
                cur = cur.next;
                j++;
            }

            cur.value = value;
        }
    }

    public T? Head() => this.head.value;

    public List Tail() => new List(){
        length = this.Length - 1,
        head = this.head.next
    };

    public int Find(T val)
    {
        Node cur = this.head;
        for (int i = 0; i < this.Length; i++)
        {
            if (cur.value.CompareTo(val) == 0)
                return i;

            cur = cur.next;
        }

        return -1;
    }

    public override string ToString()
    {
        if (this.head is null) return "[]";

        StringBuilder sb = new("[");

        Node current = this.head;

        while (current.next is not null)
        {
            sb.AppendFormat("{0}, ", current.value);
            current = current.next;
        }

        return sb.AppendFormat("{0}]", current.value).ToString();
    }
}
