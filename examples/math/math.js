
var nf = nf || {};
nf.math = {};
nf.util = {};

nf.list = function(type, cnt) {
    //TODO
}

nf.it = function() {}
nf.what = function() {}

nf.math.expression = function(s) {
    return s.split("").join('*');
}

nf.math.to_number = function(s) {
    return Number(s);
}

nf.math.decimal = function(s) {
    s = s.toString();
    var n = Number(s);
    return n / Math.pow(10, s.length);
}
