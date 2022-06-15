import _debounce  from "lodash/debounce";
import _find      from "lodash/find";
import _forEach   from "lodash/forEach";
import _get       from "lodash/get";
import _includes  from "lodash/includes";
import _indexOf   from "lodash/indexOf";
import _isNumber  from "lodash/isNumber";
import _isString  from "lodash/isString";
import _random    from "lodash/random";
import _replace   from "lodash/replace";
import _round     from "lodash/round";
import _slice     from "lodash/slice";
import _split     from "lodash/split";
import _values    from "lodash/values";


const _     = { };
_.debounce  = _debounce;
_.find      = _find;
_.forEach   = _forEach;
_.get       = _get;
_.includes  = _includes;
_.indexOf   = _indexOf;
_.isNumber  = _isNumber;
_.isString  = _isString;
_.random    = _random;
_.replace   = _replace;
_.round     = _round;
_.slice     = _slice;
_.split     = _split;
_.values    = _values;

// Custom build for lodash
// Cherry-pick methods for smaller webpack bundles
// See https://github.com/lodash/lodash#installation
export { _debounce };
export { _find };
export { _forEach };
export { _get };
export { _includes };
export { _indexOf };
export { _isNumber };
export { _isString };
export { _random };
export { _round };
export { _slice };
export { _split };
export { _replace };
export { _values };
export default _;
