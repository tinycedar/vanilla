package test;

public class MathTest {

    public static void main(String[] args){
        System.out.println("Hello world !");
        int sum = sum(10);
        System.out.println(sum);
    }

    public static int sum(int len) {
        int sum = 0;
        for (int i = 1; i <= len; i++) {
            sum += i;
        }
        return sum;
    }
}
