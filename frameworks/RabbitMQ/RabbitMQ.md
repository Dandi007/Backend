# RabbitMQ

基本概念:

* Producer(Message的生产端)

* Consumer(Message的消费端)

* Exchange(接受Producer产生的Message,将其route到合适的Queue)

* Queue(接受从Exchange route过来的Message,用于给Consumer消费)

* Connection(一条和MQ的链接,通常代表一条实际的TCP链接)

* Channel(从Connection构造的channel,由于通信,比TCP链接更加轻量)

基本流程为,Producer不断地生产数据,发送到exchange.

Consumer则通过Queue来消费数据,返回ACK.

具体的细节在于Exchange的类型,Queue和Exchange的binding和route等.

not much to say.大多是实际编程的detail.
