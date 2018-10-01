'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

function _toConsumableArray(arr) { if (Array.isArray(arr)) { for (var i = 0, arr2 = Array(arr.length); i < arr.length; i++) { arr2[i] = arr[i]; } return arr2; } else { return Array.from(arr); } }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

var Transcriptor = function () {
  function Transcriptor() {
    _classCallCheck(this, Transcriptor);

    this.complements = {
      'G': 'C',
      'C': 'G',
      'T': 'A',
      'A': 'U'
    };
  }

  _createClass(Transcriptor, [{
    key: 'toRna',
    value: function toRna(strand) {
      var _this = this;

      return [].concat(_toConsumableArray(strand)).map(function (c) {
        if (_this.complements[c] === undefined) {
          throw new Error('Invalid input DNA.');
        }
        return _this.complements[c];
      }).join('');
    }
  }]);

  return Transcriptor;
}();

exports.default = Transcriptor;