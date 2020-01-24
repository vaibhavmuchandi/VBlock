var express = require('express');
var router = express.Router();
var bodyParser = require('body-parser');

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

var vlm = require("./FabricHelper")


// Request LC
router.post('/createCar', function (req, res) {

    vlm.createCar(req, res);

});

// Issue LC
router.post('/transferCar', function (req, res) {

    vlm.transferCar(req, res);

});

// Accept LC
router.post('/sellnRegisterCar', function (req, res) {

    vlm.sellnRegisterCar(req, res);

});

// Get LC
router.post('/getCar', function (req, res) {

    vlm.getCar(req, res);

});

//Get LC history
router.post('/scrapCar', function (req, res) {

    vlm.scrapCar(req, res);

});

router.post('/getCarHistory', function (req, res) {
    vlm.getCarHistory(req, res);
});

router.post('/issueChallan', function (req, res) {
    vlm.issueChallan(req, res);
});

router.post('/payChallan', function (req, res) {
    vlm.payChallan(req, res);
});

router.post('/registerClaim', function (req, res) {
    vlm.registerClaim(req, res);
});

router.post('/payCarLoan', function (req, res) {
    vlm.payCarLoan(req, res);
});

router.post('/getCarByNumber', function (req, res) {
    vlm.getCarByNumber(req, res);
});


module.exports = router;