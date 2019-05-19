import org.openqa.selenium.By;

import static com.codeborne.selenide.CollectionCondition.size;
import static com.codeborne.selenide.Condition.text;
import static com.codeborne.selenide.Condition.value;
import static com.codeborne.selenide.Selectors.byClassName;
import static com.codeborne.selenide.Selectors.byXpath;
import static com.codeborne.selenide.Selenide.*;

import org.testng.annotations.*;
import org.testng.annotations.AfterTest;
import org.testng.annotations.BeforeTest;
import org.testng.annotations.Test;

public class MyTest {
    @BeforeTest
    public void initialTests(){
        System.out.print("Test start");
    }
    @Test
    public void checkButtonPresent() {
        open("http://localhost:8181/");
        $(byXpath("/html/body/form/input[2]")).shouldHave(value("Send"));
    }
    @Test
    @Parameters({"data"})
    public void checkDataTransfer(String data) {
        open("http://localhost:8181/");
        $(byXpath("/html/body/form/input[1]")).setValue(data).pressEnter();
        $(byXpath("/html/body/pre")).shouldHave(text("Data is " + data));
    }
    @AfterTest
    public void determTests(){
        System.out.print("Test finished");
    }

}

