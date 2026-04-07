# Tarea 1 Compiladores


Link a github: https://github.com/1andre19/Tarea-1-Compiladores/edit/main/README.md <br>
En este programa se implementaron 3 estructuras (stack, queue, hash map) desde 0 usando el lenguaje de programación Go.
Para correr el programa se necesita tener instalado Go y correr el siguiente comando en el root del proyecto
```
go run .
```

## Test Cases
Se verificó que el **stack** se inicializara vacío con IsEmpty() = true y Size() = 0. 
Se insertaron tres valores 10, 20, 30 y se verificó que Peek() devuelve el último elemento pusheado (30). 
Se hizo Pop() y se confirmó que el nuevo tope es 20. 
Finalmente, se hizo Pop() hasta vaciar la pila.<br>

Se verificó que el **queue** inicia vacío. Se encolaron tres valores (1, 2, 3) y se comprobó que el tamaño es 3. 
Se hizo Dequeue() y se verificó que el tamaño decrementó correctamente. 
Se vació la cola por completo y se confirmó que al hacer Dequeue() de todos los elementos la **queue** esta vacía. <br>

Se insertaron tres key-value pairs (alice=30, bob=25, carol=28) y se verificó que Get() retorna el valor correcto. 
Se sobreescribió bob con el valor 99 y  Get("bob") retorna el nuevo valor. 
Se intentó obtener una clave inexistente (dave) y se verificó que ok = false. 
Por último, se eliminó alice con Delete() y se confirmó que ya no existe en el mapa.

