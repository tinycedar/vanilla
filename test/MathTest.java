package test;

public class MathTest {

    public static void main(String[] args){
        System.out.println("Hello world !");
        int sum = sum(10);
        System.out.println(sum);
        System.out.println(addFloat());
    }

    public static int sum(int len) {
        int sum = 0;
        for (int i = 1; i <= len; i++) {
            sum += i;
        }
        return sum;
    }

    public static float addFloat() {
        float a = 1.2f;
        float b = 2.3f;
        return a + b;
    }
}
