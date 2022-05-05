# Manual de Usuario

### Requisitos

1. Internet
2. Postman o Insomnia

### Url

Para poder hacer uso de la practica con la arquitectura es necesario correr la siguiente Url a traves de Postman o Insomnia:

![](images/postman.png)


```
http://34.135.238.242:30347/api/exec-game
```

Esto a traves de una peticion de tipo `Post`:


![](images/log.png)

Al hacer la peticion deberiamos ser capaz de ver un mensaje asi:


![](images/log1.png)


Esto significa que nuestro Log fue enviado exitosamente


## Comprobar en la base de datos


Para comprobar que la informacion fue enviada de forma correcta podemos acceder a la base de datos de MongoDB para visualizar el log almacenado:

![](images/mongo.png)


La imagen de arriba muestra como es que los logs se deberian almacenar.