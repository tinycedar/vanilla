package test;

public class ComparisonTest {

    public static void main(String[] args){

        int a = 1;
        if (a == 0) {
            System.out.println("a == 0");
        }
        if (a != 0) {
            System.out.println("a != 0");
        }
        if (a > 0) {
            System.out.println("a > 0");
        }
        if (a >= 0) {
            System.out.println("a >= 0");
        }
        if (a < 0) {
            System.out.println("a < 0");
        }
        if (a <= 0) {
            System.out.println("a <= 0");
        }
        rate(50);
        rate(75);
        rate(90);
    }

    public static void rate(int score) {
        if (score < 60) {
            System.out.println("不及格 !");
        } else if (score < 75) {
            System.out.println("及格 !");
        } else if (score < 85) {
            System.out.println("良好 !");
        } else {
            System.out.println("优秀 !");
        }
    }
}