import rclpy
from rclpy.node import Node

from std_msgs.msg import String


class MinPubSub(Node):

    def __init__(self):
        super().__init__('minimal_node')
        self.publisher_ = self.create_publisher(String, '/cat_topic', 10)
        timer_period = 2.0  # seconds
        self.timer = self.create_timer(timer_period, self.timer_callback)
        self.i = 1

        self.subscription = self.create_subscription(
            String,
            '/dog_topic',
            self.listener_callback,
            10)
        self.subscription  # prevent unused variable warning

    def timer_callback(self):
        msg = String()
        msg.data = 'Meow %d' % self.i
        self.publisher_.publish(msg)
        self.get_logger().info('Publishing: "%s"' % msg.data)
        self.i += 1

    def listener_callback(self, msg):
        self.get_logger().info('Received msg: "%s"' % msg.data)


def main(args=None):
    rclpy.init(args=args)

    pubsub = MinPubSub()

    rclpy.spin(pubsub)

    # Destroy the node explicitly
    # (optional - otherwise it will be done automatically
    # when the garbage collector destroys the node object)
    pubsub.destroy_node()
    rclpy.shutdown()


if __name__ == '__main__':
    main()
