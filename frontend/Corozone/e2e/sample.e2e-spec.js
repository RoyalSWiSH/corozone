const nsAppium = require("nativescript-dev-appium");
const assert = require("chai").assert;
const addContext = require('mochawesome/addContext');

describe("sample scenario", () => {
    let driver;

    before(async function() {
        nsAppium.nsCapabilities.testReporter.context = this; 
        driver = await nsAppium.createDriver();
    });

    after(async function () {
        await driver.quit();
        console.log("Quit driver!");
    });

    afterEach(async function () {
        if (this.currentTest.state === "failed") {
            await driver.logTestArtifacts(this.currentTest.title);
        }
    });

    it("Ask if passowrd forgotten", async function () {
        const label = await driver.findElementByText("Forgot your Password?", "contains");
        assert.isTrue(await label.isDisplayed());
    });
    it("Show login Button", async function () {
        const label = await driver.findElementByText("Log In", "contains");
        assert.isTrue(await label.isDisplayed());
    });

    it("Test Login", async function () {
       // const label = await driver.findElementByText("Log In", "contains");
        //const mailField = await driver.findElementByAccessibilityId("Email");

        const EMAIL = "asd@asd.de"
        const PASSWORD = "asdasd"
        const allFields = await driver.findElementsByClassName("android.widget.EditText");
        await allFields[0].sendKeys(EMAIL);
        await allFields[1].sendKeys(PASSWORD);
        await driver.hideDeviceKeyboard();
        //await mailField.sendKeys("2");

        const tapButton = await driver.findElementByText("Log In");
        await tapButton.click();
        await driver.wait(6000);
        //await  driver.elementByXPath('//android.widget.TextView[@text=\'Log In\']').click();
        const label = await driver.findElementByText("Reload", "contains");
        //const notconnectedmsg = await driver.findElementByText("Not connected.", "contains");
        const detailsButton = await driver.findElementByText("Details", "contains");
       // assert.isTrue(await label.isDisplayed());
       // assert.isTrue(await detailsButton.isDisplayed());
        await detailsButton.click() 
        await driver.wait(8000);
       // await detailsButton.click() 
        const acceptButton = await driver.findElementByText("Google Maps", "contains");
        assert.isTrue(await acceptButton.isDisplayed()); 
    });

    // it("Test Details", async function () {
    //     const detailsButton = await driver.findElementByText("Details", "contains");
    //     await detailsButton.click() 
    //     await driver.wait(5000);
    //     const acceptButton = await driver.findElementByText("Accept", "contains");
    //     assert.isTrue(await acceptButton.isDisplayed()); 
    // });
});