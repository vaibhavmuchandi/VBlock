var express = require('express');
var router = express.Router();
var bodyParser = require('body-parser');

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

var vlm = require("./FabricHelper")


// Create Car
router.post('/createCar', function (req, res) {

    vlm.createCar(req, res);

});

// Transfer Car to Dealer
router.post('/transferCar', function (req, res) {

    vlm.transferCar(req, res);

});

// Sell and register car to customer
router.post('/sellnRegisterCar', function (req, res) {

    vlm.sellnRegisterCar(req, res);

});

// Get car details
router.post('/getCar', function (req, res) {

    vlm.getCar(req, res);

});

//Scrap car
router.post('/scrapCar', function (req, res) {

    vlm.scrapCar(req, res);

});

//Get car history
router.post('/getCarHistory', function (req, res) {
    vlm.getCarHistory(req, res);
});

//Issue challan
router.post('/issueChallan', function (req, res) {
    vlm.issueChallan(req, res);
});

//Pay challan
router.post('/payChallan', function (req, res) {
    vlm.payChallan(req, res);
});

//Register insurance claim
router.post('/registerClaim', function (req, res) {
    vlm.registerClaim(req, res);
});

//Pay Car loan
router.post('/payCarLoan', function (req, res) {
    vlm.payCarLoan(req, res);
});

//Get car by regis no
router.post('/getCarByNumber', function (req, res) {
    vlm.getCarByNumber(req, res);
});


module.exports = router;
