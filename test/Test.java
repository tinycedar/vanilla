public class Test {

    public static void main(String[] args){
        int a = 1;
        int b = a + 2;
        System.out.println("Hello world !");
        System.out.println(b);

        staticMethod();
    }

    public static void staticMethod() {
        System.out.println(sum());
        System.out.println("========= static method invoked ! ==========");
    }

    public static int sum() {
        int sum = 0;
        for (int i = 1; i < 101; i++) {
            sum += i;
        }
        return sum;
    }
}
