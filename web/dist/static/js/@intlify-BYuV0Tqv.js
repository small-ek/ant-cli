var e=Object.defineProperty,t=Object.defineProperties,n=Object.getOwnPropertyDescriptors,r=Object.getOwnPropertySymbols,o=Object.prototype.hasOwnProperty,s=Object.prototype.propertyIsEnumerable,c=(t,n,r)=>n in t?e(t,n,{enumerable:!0,configurable:!0,writable:!0,value:r}):t[n]=r,a=(e,t)=>{for(var n in t||(t={}))o.call(t,n)&&c(e,n,t[n]);if(r)for(var n of r(t))s.call(t,n)&&c(e,n,t[n]);return e},l=(e,r)=>t(e,n(r))
/*!
  * shared v9.13.1
  * (c) 2024 kazuya kawaguchi
  * Released under the MIT License.
  */;const i="undefined"!=typeof window,u=(e,t=!1)=>t?Symbol.for(e):Symbol(e),f=(e,t,n)=>_({l:e,k:t,s:n}),_=e=>JSON.stringify(e).replace(/\u2028/g,"\\u2028").replace(/\u2029/g,"\\u2029").replace(/\u0027/g,"\\u0027"),E=e=>"number"==typeof e&&isFinite(e),d=e=>"[object Date]"===S(e),p=e=>"[object RegExp]"===S(e),m=e=>D(e)&&0===Object.keys(e).length,N=Object.assign;let L;const T=()=>L||(L="undefined"!=typeof globalThis?globalThis:"undefined"!=typeof self?self:"undefined"!=typeof window?window:"undefined"!=typeof global?global:{});function h(e){return e.replace(/</g,"&lt;").replace(/>/g,"&gt;").replace(/"/g,"&quot;").replace(/'/g,"&apos;")}const I=Object.prototype.hasOwnProperty;function O(e,t){return I.call(e,t)}const A=Array.isArray,k=e=>"function"==typeof e,y=e=>"string"==typeof e,C=e=>"boolean"==typeof e,b=e=>null!==e&&"object"==typeof e,g=e=>b(e)&&k(e.then)&&k(e.catch),P=Object.prototype.toString,S=e=>P.call(e),D=e=>{if(!b(e))return!1;const t=Object.getPrototypeOf(e);return null===t||t.constructor===Object};function U(e){let t=e;return()=>++t}function v(e,t){}const R=e=>!b(e)||A(e);function M(e,t){if(R(e)||R(t))throw new Error("Invalid value");const n=[{src:e,des:t}];for(;n.length;){const{src:e,des:t}=n.pop();Object.keys(e).forEach((r=>{R(e[r])||R(t[r])?t[r]=e[r]:n.push({src:e[r],des:t[r]})}))}}
/*!
  * message-compiler v9.13.1
  * (c) 2024 kazuya kawaguchi
  * Released under the MIT License.
  */function w(e,t,n){const r={start:e,end:t};return null!=n&&(r.source=n),r}const x=/\{([0-9a-zA-Z]+)\}/g;function F(e,...t){return 1===t.length&&W(t[0])&&(t=t[0]),t&&t.hasOwnProperty||(t={}),e.replace(x,((e,n)=>t.hasOwnProperty(n)?t[n]:""))}const Y=Object.assign,X=e=>"string"==typeof e,W=e=>null!==e&&"object"==typeof e;function $(e,t=""){return e.reduce(((e,n,r)=>0===r?e+n:e+t+n),"")}const G={USE_MODULO_SYNTAX:1,__EXTEND_POINT__:2},K={[G.USE_MODULO_SYNTAX]:"Use modulo before '{{0}}'."};const H={EXPECTED_TOKEN:1,INVALID_TOKEN_IN_PLACEHOLDER:2,UNTERMINATED_SINGLE_QUOTE_IN_PLACEHOLDER:3,UNKNOWN_ESCAPE_SEQUENCE:4,INVALID_UNICODE_ESCAPE_SEQUENCE:5,UNBALANCED_CLOSING_BRACE:6,UNTERMINATED_CLOSING_BRACE:7,EMPTY_PLACEHOLDER:8,NOT_ALLOW_NEST_PLACEHOLDER:9,INVALID_LINKED_FORMAT:10,MUST_HAVE_MESSAGES_IN_PLURAL:11,UNEXPECTED_EMPTY_LINKED_MODIFIER:12,UNEXPECTED_EMPTY_LINKED_KEY:13,UNEXPECTED_LEXICAL_ANALYSIS:14,UNHANDLED_CODEGEN_NODE_TYPE:15,UNHANDLED_MINIFIER_NODE_TYPE:16,__EXTEND_POINT__:17},V={[H.EXPECTED_TOKEN]:"Expected token: '{0}'",[H.INVALID_TOKEN_IN_PLACEHOLDER]:"Invalid token in placeholder: '{0}'",[H.UNTERMINATED_SINGLE_QUOTE_IN_PLACEHOLDER]:"Unterminated single quote in placeholder",[H.UNKNOWN_ESCAPE_SEQUENCE]:"Unknown escape sequence: \\{0}",[H.INVALID_UNICODE_ESCAPE_SEQUENCE]:"Invalid unicode escape sequence: {0}",[H.UNBALANCED_CLOSING_BRACE]:"Unbalanced closing brace",[H.UNTERMINATED_CLOSING_BRACE]:"Unterminated closing brace",[H.EMPTY_PLACEHOLDER]:"Empty placeholder",[H.NOT_ALLOW_NEST_PLACEHOLDER]:"Not allowed nest placeholder",[H.INVALID_LINKED_FORMAT]:"Invalid linked format",[H.MUST_HAVE_MESSAGES_IN_PLURAL]:"Plural must have messages",[H.UNEXPECTED_EMPTY_LINKED_MODIFIER]:"Unexpected empty linked modifier",[H.UNEXPECTED_EMPTY_LINKED_KEY]:"Unexpected empty linked key",[H.UNEXPECTED_LEXICAL_ANALYSIS]:"Unexpected lexical analysis in token: '{0}'",[H.UNHANDLED_CODEGEN_NODE_TYPE]:"unhandled codegen node type: '{0}'",[H.UNHANDLED_MINIFIER_NODE_TYPE]:"unhandled mimifier node type: '{0}'"};function j(e,t,n={}){const{domain:r,messages:o,args:s}=n,c=F((o||V)[e]||"",...s||[]),a=new SyntaxError(String(c));return a.code=e,t&&(a.location=t),a.domain=r,a}function B(e){throw e}const z=" ",J="\r",Q="\n",q=String.fromCharCode(8232),Z=String.fromCharCode(8233);function ee(e){const t=e;let n=0,r=1,o=1,s=0;const c=e=>t[e]===J&&t[e+1]===Q,a=e=>t[e]===Z,l=e=>t[e]===q,i=e=>c(e)||(e=>t[e]===Q)(e)||a(e)||l(e),u=e=>c(e)||a(e)||l(e)?Q:t[e];function f(){return s=0,i(n)&&(r++,o=0),c(n)&&n++,n++,o++,t[n]}return{index:()=>n,line:()=>r,column:()=>o,peekOffset:()=>s,charAt:u,currentChar:()=>u(n),currentPeek:()=>u(n+s),next:f,peek:function(){return c(n+s)&&s++,s++,t[n+s]},reset:function(){n=0,r=1,o=1,s=0},resetPeek:function(e=0){s=e},skipToPeek:function(){const e=n+s;for(;e!==n;)f();s=0}}}const te=void 0,ne="'",re="tokenizer";function oe(e,t={}){const n=!1!==t.location,r=ee(e),o=()=>r.index(),s=()=>{return e=r.line(),t=r.column(),n=r.index(),{line:e,column:t,offset:n};var e,t,n},c=s(),a=o(),l={currentType:14,offset:a,startLoc:c,endLoc:c,lastType:14,lastOffset:a,lastStartLoc:c,lastEndLoc:c,braceNest:0,inLinked:!1,text:""},i=()=>l,{onError:u}=t;function f(e,t,r,...o){const s=i();if(t.column+=r,t.offset+=r,u){const r=j(e,n?w(s.startLoc,t):null,{domain:re,args:o});u(r)}}function _(e,t,r){e.endLoc=s(),e.currentType=t;const o={type:t};return n&&(o.loc=w(e.startLoc,e.endLoc)),null!=r&&(o.value=r),o}const E=e=>_(e,14);function d(e,t){return e.currentChar()===t?(e.next(),t):(f(H.EXPECTED_TOKEN,s(),0,t),"")}function p(e){let t="";for(;e.currentPeek()===z||e.currentPeek()===Q;)t+=e.currentPeek(),e.peek();return t}function m(e){const t=p(e);return e.skipToPeek(),t}function N(e){if(e===te)return!1;const t=e.charCodeAt(0);return t>=97&&t<=122||t>=65&&t<=90||95===t}function L(e,t){const{currentType:n}=t;if(2!==n)return!1;p(e);const r=function(e){if(e===te)return!1;const t=e.charCodeAt(0);return t>=48&&t<=57}("-"===e.currentPeek()?e.peek():e.currentPeek());return e.resetPeek(),r}function T(e){p(e);const t="|"===e.currentPeek();return e.resetPeek(),t}function h(e,t=!0){const n=(t=!1,r="",o=!1)=>{const s=e.currentPeek();return"{"===s?"%"!==r&&t:"@"!==s&&s?"%"===s?(e.peek(),n(t,"%",!0)):"|"===s?!("%"!==r&&!o)||!(r===z||r===Q):s===z?(e.peek(),n(!0,z,o)):s!==Q||(e.peek(),n(!0,Q,o)):"%"===r||t},r=n();return t&&e.resetPeek(),r}function I(e,t){const n=e.currentChar();return n===te?te:t(n)?(e.next(),n):null}function O(e){const t=e.charCodeAt(0);return t>=97&&t<=122||t>=65&&t<=90||t>=48&&t<=57||95===t||36===t}function A(e){return I(e,O)}function k(e){const t=e.charCodeAt(0);return t>=97&&t<=122||t>=65&&t<=90||t>=48&&t<=57||95===t||36===t||45===t}function y(e){return I(e,k)}function C(e){const t=e.charCodeAt(0);return t>=48&&t<=57}function b(e){return I(e,C)}function g(e){const t=e.charCodeAt(0);return t>=48&&t<=57||t>=65&&t<=70||t>=97&&t<=102}function P(e){return I(e,g)}function S(e){let t="",n="";for(;t=b(e);)n+=t;return n}function D(e){let t="";for(;;){const n=e.currentChar();if("{"===n||"}"===n||"@"===n||"|"===n||!n)break;if("%"===n){if(!h(e))break;t+=n,e.next()}else if(n===z||n===Q)if(h(e))t+=n,e.next();else{if(T(e))break;t+=n,e.next()}else t+=n,e.next()}return t}function U(e){return e!==ne&&e!==Q}function v(e){const t=e.currentChar();switch(t){case"\\":case"'":return e.next(),`\\${t}`;case"u":return R(e,t,4);case"U":return R(e,t,6);default:return f(H.UNKNOWN_ESCAPE_SEQUENCE,s(),0,t),""}}function R(e,t,n){d(e,t);let r="";for(let o=0;o<n;o++){const n=P(e);if(!n){f(H.INVALID_UNICODE_ESCAPE_SEQUENCE,s(),0,`\\${t}${r}${e.currentChar()}`);break}r+=n}return`\\${t}${r}`}function M(e){return"{"!==e&&"}"!==e&&e!==z&&e!==Q}function x(e){m(e);const t=d(e,"|");return m(e),t}function F(e,t){let n=null;switch(e.currentChar()){case"{":return t.braceNest>=1&&f(H.NOT_ALLOW_NEST_PLACEHOLDER,s(),0),e.next(),n=_(t,2,"{"),m(e),t.braceNest++,n;case"}":return t.braceNest>0&&2===t.currentType&&f(H.EMPTY_PLACEHOLDER,s(),0),e.next(),n=_(t,3,"}"),t.braceNest--,t.braceNest>0&&m(e),t.inLinked&&0===t.braceNest&&(t.inLinked=!1),n;case"@":return t.braceNest>0&&f(H.UNTERMINATED_CLOSING_BRACE,s(),0),n=Y(e,t)||E(t),t.braceNest=0,n;default:{let r=!0,o=!0,c=!0;if(T(e))return t.braceNest>0&&f(H.UNTERMINATED_CLOSING_BRACE,s(),0),n=_(t,1,x(e)),t.braceNest=0,t.inLinked=!1,n;if(t.braceNest>0&&(5===t.currentType||6===t.currentType||7===t.currentType))return f(H.UNTERMINATED_CLOSING_BRACE,s(),0),t.braceNest=0,X(e,t);if(r=function(e,t){const{currentType:n}=t;if(2!==n)return!1;p(e);const r=N(e.currentPeek());return e.resetPeek(),r}(e,t))return n=_(t,5,function(e){m(e);let t="",n="";for(;t=y(e);)n+=t;return e.currentChar()===te&&f(H.UNTERMINATED_CLOSING_BRACE,s(),0),n}(e)),m(e),n;if(o=L(e,t))return n=_(t,6,function(e){m(e);let t="";return"-"===e.currentChar()?(e.next(),t+=`-${S(e)}`):t+=S(e),e.currentChar()===te&&f(H.UNTERMINATED_CLOSING_BRACE,s(),0),t}(e)),m(e),n;if(c=function(e,t){const{currentType:n}=t;if(2!==n)return!1;p(e);const r=e.currentPeek()===ne;return e.resetPeek(),r}(e,t))return n=_(t,7,function(e){m(e),d(e,"'");let t="",n="";for(;t=I(e,U);)n+="\\"===t?v(e):t;const r=e.currentChar();return r===Q||r===te?(f(H.UNTERMINATED_SINGLE_QUOTE_IN_PLACEHOLDER,s(),0),r===Q&&(e.next(),d(e,"'")),n):(d(e,"'"),n)}(e)),m(e),n;if(!r&&!o&&!c)return n=_(t,13,function(e){m(e);let t="",n="";for(;t=I(e,M);)n+=t;return n}(e)),f(H.INVALID_TOKEN_IN_PLACEHOLDER,s(),0,n.value),m(e),n;break}}return n}function Y(e,t){const{currentType:n}=t;let r=null;const o=e.currentChar();switch(8!==n&&9!==n&&12!==n&&10!==n||o!==Q&&o!==z||f(H.INVALID_LINKED_FORMAT,s(),0),o){case"@":return e.next(),r=_(t,8,"@"),t.inLinked=!0,r;case".":return m(e),e.next(),_(t,9,".");case":":return m(e),e.next(),_(t,10,":");default:return T(e)?(r=_(t,1,x(e)),t.braceNest=0,t.inLinked=!1,r):function(e,t){const{currentType:n}=t;if(8!==n)return!1;p(e);const r="."===e.currentPeek();return e.resetPeek(),r}(e,t)||function(e,t){const{currentType:n}=t;if(8!==n&&12!==n)return!1;p(e);const r=":"===e.currentPeek();return e.resetPeek(),r}(e,t)?(m(e),Y(e,t)):function(e,t){const{currentType:n}=t;if(9!==n)return!1;p(e);const r=N(e.currentPeek());return e.resetPeek(),r}(e,t)?(m(e),_(t,12,function(e){let t="",n="";for(;t=A(e);)n+=t;return n}(e))):function(e,t){const{currentType:n}=t;if(10!==n)return!1;const r=()=>{const t=e.currentPeek();return"{"===t?N(e.peek()):!("@"===t||"%"===t||"|"===t||":"===t||"."===t||t===z||!t)&&(t===Q?(e.peek(),r()):h(e,!1))},o=r();return e.resetPeek(),o}(e,t)?(m(e),"{"===o?F(e,t)||r:_(t,11,function(e){const t=n=>{const r=e.currentChar();return"{"!==r&&"%"!==r&&"@"!==r&&"|"!==r&&"("!==r&&")"!==r&&r?r===z?n:(n+=r,e.next(),t(n)):n};return t("")}(e))):(8===n&&f(H.INVALID_LINKED_FORMAT,s(),0),t.braceNest=0,t.inLinked=!1,X(e,t))}}function X(e,t){let n={type:14};if(t.braceNest>0)return F(e,t)||E(t);if(t.inLinked)return Y(e,t)||E(t);switch(e.currentChar()){case"{":return F(e,t)||E(t);case"}":return f(H.UNBALANCED_CLOSING_BRACE,s(),0),e.next(),_(t,3,"}");case"@":return Y(e,t)||E(t);default:{if(T(e))return n=_(t,1,x(e)),t.braceNest=0,t.inLinked=!1,n;const{isModulo:r,hasSpace:o}=function(e){const t=p(e),n="%"===e.currentPeek()&&"{"===e.peek();return e.resetPeek(),{isModulo:n,hasSpace:t.length>0}}(e);if(r)return o?_(t,0,D(e)):_(t,4,function(e){m(e);const t=e.currentChar();return"%"!==t&&f(H.EXPECTED_TOKEN,s(),0,t),e.next(),"%"}(e));if(h(e))return _(t,0,D(e));break}}return n}return{nextToken:function(){const{currentType:e,offset:t,startLoc:n,endLoc:c}=l;return l.lastType=e,l.lastOffset=t,l.lastStartLoc=n,l.lastEndLoc=c,l.offset=o(),l.startLoc=s(),r.currentChar()===te?_(l,14):X(r,l)},currentOffset:o,currentPosition:s,context:i}}const se="parser",ce=/(?:\\\\|\\'|\\u([0-9a-fA-F]{4})|\\U([0-9a-fA-F]{6}))/g;function ae(e,t,n){switch(e){case"\\\\":return"\\";case"\\'":return"'";default:{const e=parseInt(t||n,16);return e<=55295||e>=57344?String.fromCodePoint(e):"�"}}}function le(e={}){const t=!1!==e.location,{onError:n,onWarn:r}=e;function o(e,r,o,s,...c){const a=e.currentPosition();if(a.offset+=s,a.column+=s,n){const e=j(r,t?w(o,a):null,{domain:se,args:c});n(e)}}function s(e,n,o,s,...c){const a=e.currentPosition();if(a.offset+=s,a.column+=s,r){const e=t?w(o,a):null;r(function(e,t,...n){const r=F(K[e]||"",...n||[]),o={message:String(r),code:e};return t&&(o.location=t),o}(n,e,c))}}function c(e,n,r){const o={type:e};return t&&(o.start=n,o.end=n,o.loc={start:r,end:r}),o}function a(e,n,r,o){o&&(e.type=o),t&&(e.end=n,e.loc&&(e.loc.end=r))}function l(e,t){const n=e.context(),r=c(3,n.offset,n.startLoc);return r.value=t,a(r,e.currentOffset(),e.currentPosition()),r}function i(e,t){const n=e.context(),{lastOffset:r,lastStartLoc:o}=n,s=c(5,r,o);return s.index=parseInt(t,10),e.nextToken(),a(s,e.currentOffset(),e.currentPosition()),s}function u(e,t,n){const r=e.context(),{lastOffset:o,lastStartLoc:s}=r,l=c(4,o,s);return l.key=t,!0===n&&(l.modulo=!0),e.nextToken(),a(l,e.currentOffset(),e.currentPosition()),l}function f(e,t){const n=e.context(),{lastOffset:r,lastStartLoc:o}=n,s=c(9,r,o);return s.value=t.replace(ce,ae),e.nextToken(),a(s,e.currentOffset(),e.currentPosition()),s}function _(e){const t=e.context(),n=c(6,t.offset,t.startLoc);let r=e.nextToken();if(9===r.type){const t=function(e){const t=e.nextToken(),n=e.context(),{lastOffset:r,lastStartLoc:s}=n,l=c(8,r,s);return 12!==t.type?(o(e,H.UNEXPECTED_EMPTY_LINKED_MODIFIER,n.lastStartLoc,0),l.value="",a(l,r,s),{nextConsumeToken:t,node:l}):(null==t.value&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,n.lastStartLoc,0,ie(t)),l.value=t.value||"",a(l,e.currentOffset(),e.currentPosition()),{node:l})}(e);n.modifier=t.node,r=t.nextConsumeToken||e.nextToken()}switch(10!==r.type&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,t.lastStartLoc,0,ie(r)),r=e.nextToken(),2===r.type&&(r=e.nextToken()),r.type){case 11:null==r.value&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,t.lastStartLoc,0,ie(r)),n.key=function(e,t){const n=e.context(),r=c(7,n.offset,n.startLoc);return r.value=t,a(r,e.currentOffset(),e.currentPosition()),r}(e,r.value||"");break;case 5:null==r.value&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,t.lastStartLoc,0,ie(r)),n.key=u(e,r.value||"");break;case 6:null==r.value&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,t.lastStartLoc,0,ie(r)),n.key=i(e,r.value||"");break;case 7:null==r.value&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,t.lastStartLoc,0,ie(r)),n.key=f(e,r.value||"");break;default:{o(e,H.UNEXPECTED_EMPTY_LINKED_KEY,t.lastStartLoc,0);const s=e.context(),l=c(7,s.offset,s.startLoc);return l.value="",a(l,s.offset,s.startLoc),n.key=l,a(n,s.offset,s.startLoc),{nextConsumeToken:r,node:n}}}return a(n,e.currentOffset(),e.currentPosition()),{node:n}}function E(e){const t=e.context(),n=c(2,1===t.currentType?e.currentOffset():t.offset,1===t.currentType?t.endLoc:t.startLoc);n.items=[];let r=null,E=null;do{const c=r||e.nextToken();switch(r=null,c.type){case 0:null==c.value&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,t.lastStartLoc,0,ie(c)),n.items.push(l(e,c.value||""));break;case 6:null==c.value&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,t.lastStartLoc,0,ie(c)),n.items.push(i(e,c.value||""));break;case 4:E=!0;break;case 5:null==c.value&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,t.lastStartLoc,0,ie(c)),n.items.push(u(e,c.value||"",!!E)),E&&(s(e,G.USE_MODULO_SYNTAX,t.lastStartLoc,0,ie(c)),E=null);break;case 7:null==c.value&&o(e,H.UNEXPECTED_LEXICAL_ANALYSIS,t.lastStartLoc,0,ie(c)),n.items.push(f(e,c.value||""));break;case 8:{const t=_(e);n.items.push(t.node),r=t.nextConsumeToken||null;break}}}while(14!==t.currentType&&1!==t.currentType);return a(n,1===t.currentType?t.lastOffset:e.currentOffset(),1===t.currentType?t.lastEndLoc:e.currentPosition()),n}function d(e){const t=e.context(),{offset:n,startLoc:r}=t,s=E(e);return 14===t.currentType?s:function(e,t,n,r){const s=e.context();let l=0===r.items.length;const i=c(1,t,n);i.cases=[],i.cases.push(r);do{const t=E(e);l||(l=0===t.items.length),i.cases.push(t)}while(14!==s.currentType);return l&&o(e,H.MUST_HAVE_MESSAGES_IN_PLURAL,n,0),a(i,e.currentOffset(),e.currentPosition()),i}(e,n,r,s)}return{parse:function(n){const r=oe(n,Y({},e)),s=r.context(),l=c(0,s.offset,s.startLoc);return t&&l.loc&&(l.loc.source=n),l.body=d(r),e.onCacheKey&&(l.cacheKey=e.onCacheKey(n)),14!==s.currentType&&o(r,H.UNEXPECTED_LEXICAL_ANALYSIS,s.lastStartLoc,0,n[s.offset]||""),a(l,r.currentOffset(),r.currentPosition()),l}}}function ie(e){if(14===e.type)return"EOF";const t=(e.value||"").replace(/\r?\n/gu,"\\n");return t.length>10?t.slice(0,9)+"…":t}function ue(e,t){for(let n=0;n<e.length;n++)fe(e[n],t)}function fe(e,t){switch(e.type){case 1:ue(e.cases,t),t.helper("plural");break;case 2:ue(e.items,t);break;case 6:fe(e.key,t),t.helper("linked"),t.helper("type");break;case 5:t.helper("interpolate"),t.helper("list");break;case 4:t.helper("interpolate"),t.helper("named")}}function _e(e,t={}){const n=function(e,t={}){const n={ast:e,helpers:new Set};return{context:()=>n,helper:e=>(n.helpers.add(e),e)}}(e);n.helper("normalize"),e.body&&fe(e.body,n);const r=n.context();e.helpers=Array.from(r.helpers)}function Ee(e){if(1===e.items.length){const t=e.items[0];3!==t.type&&9!==t.type||(e.static=t.value,delete t.value)}else{const t=[];for(let n=0;n<e.items.length;n++){const r=e.items[n];if(3!==r.type&&9!==r.type)break;if(null==r.value)break;t.push(r.value)}if(t.length===e.items.length){e.static=$(t);for(let t=0;t<e.items.length;t++){const n=e.items[t];3!==n.type&&9!==n.type||delete n.value}}}}const de="minifier";function pe(e){switch(e.t=e.type,e.type){case 0:{const t=e;pe(t.body),t.b=t.body,delete t.body;break}case 1:{const t=e,n=t.cases;for(let e=0;e<n.length;e++)pe(n[e]);t.c=n,delete t.cases;break}case 2:{const t=e,n=t.items;for(let e=0;e<n.length;e++)pe(n[e]);t.i=n,delete t.items,t.static&&(t.s=t.static,delete t.static);break}case 3:case 9:case 8:case 7:{const t=e;t.value&&(t.v=t.value,delete t.value);break}case 6:{const t=e;pe(t.key),t.k=t.key,delete t.key,t.modifier&&(pe(t.modifier),t.m=t.modifier,delete t.modifier);break}case 5:{const t=e;t.i=t.index,delete t.index;break}case 4:{const t=e;t.k=t.key,delete t.key;break}default:throw j(H.UNHANDLED_MINIFIER_NODE_TYPE,null,{domain:de,args:[e.type]})}delete e.type}const me="parser";function Ne(e,t){const{helper:n}=e;switch(t.type){case 0:!function(e,t){t.body?Ne(e,t.body):e.push("null")}(e,t);break;case 1:!function(e,t){const{helper:n,needIndent:r}=e;if(t.cases.length>1){e.push(`${n("plural")}([`),e.indent(r());const o=t.cases.length;for(let n=0;n<o&&(Ne(e,t.cases[n]),n!==o-1);n++)e.push(", ");e.deindent(r()),e.push("])")}}(e,t);break;case 2:!function(e,t){const{helper:n,needIndent:r}=e;e.push(`${n("normalize")}([`),e.indent(r());const o=t.items.length;for(let s=0;s<o&&(Ne(e,t.items[s]),s!==o-1);s++)e.push(", ");e.deindent(r()),e.push("])")}(e,t);break;case 6:!function(e,t){const{helper:n}=e;e.push(`${n("linked")}(`),Ne(e,t.key),t.modifier?(e.push(", "),Ne(e,t.modifier),e.push(", _type")):e.push(", undefined, _type"),e.push(")")}(e,t);break;case 8:case 7:case 9:case 3:e.push(JSON.stringify(t.value),t);break;case 5:e.push(`${n("interpolate")}(${n("list")}(${t.index}))`,t);break;case 4:e.push(`${n("interpolate")}(${n("named")}(${JSON.stringify(t.key)}))`,t);break;default:throw j(H.UNHANDLED_CODEGEN_NODE_TYPE,null,{domain:me,args:[t.type]})}}const Le=(e,t={})=>{const n=X(t.mode)?t.mode:"normal",r=X(t.filename)?t.filename:"message.intl",o=!!t.sourceMap,s=null!=t.breakLineCode?t.breakLineCode:"arrow"===n?";":"\n",c=t.needIndent?t.needIndent:"arrow"!==n,a=e.helpers||[],l=function(e,t){const{sourceMap:n,filename:r,breakLineCode:o,needIndent:s}=t,c=!1!==t.location,a={filename:r,code:"",column:1,line:1,offset:0,map:void 0,breakLineCode:o,needIndent:s,indentLevel:0};function l(e,t){a.code+=e}function i(e,t=!0){const n=t?o:"";l(s?n+"  ".repeat(e):n)}return c&&e.loc&&(a.source=e.loc.source),{context:()=>a,push:l,indent:function(e=!0){const t=++a.indentLevel;e&&i(t)},deindent:function(e=!0){const t=--a.indentLevel;e&&i(t)},newline:function(){i(a.indentLevel)},helper:e=>`_${e}`,needIndent:()=>a.needIndent}}(e,{mode:n,filename:r,sourceMap:o,breakLineCode:s,needIndent:c});l.push("normal"===n?"function __msg__ (ctx) {":"(ctx) => {"),l.indent(c),a.length>0&&(l.push(`const { ${$(a.map((e=>`${e}: _${e}`)),", ")} } = ctx`),l.newline()),l.push("return "),Ne(l,e),l.deindent(c),l.push("}"),delete e.helpers;const{code:i,map:u}=l.context();return{ast:e,code:i,map:u?u.toJSON():void 0}};function Te(e,t={}){const n=Y({},t),r=!!n.jit,o=!!n.minify,s=null==n.optimize||n.optimize,c=le(n).parse(e);return r?(s&&function(e){const t=e.body;2===t.type?Ee(t):t.cases.forEach((e=>Ee(e)))}(c),o&&pe(c),{ast:c,code:""}):(_e(c,n),Le(c,n))}
/*!
  * core-base v9.13.1
  * (c) 2024 kazuya kawaguchi
  * Released under the MIT License.
  */const he=[];he[0]={w:[0],i:[3,0],"[":[4],o:[7]},he[1]={w:[1],".":[2],"[":[4],o:[7]},he[2]={w:[2],i:[3,0],0:[3,0]},he[3]={i:[3,0],0:[3,0],w:[1,1],".":[2,1],"[":[4,1],o:[7,1]},he[4]={"'":[5,0],'"':[6,0],"[":[4,2],"]":[1,3],o:8,l:[4,0]},he[5]={"'":[4,0],o:8,l:[5,0]},he[6]={'"':[4,0],o:8,l:[6,0]};const Ie=/^\s?(?:true|false|-?[\d.]+|'[^']*'|"[^"]*")\s?$/;function Oe(e){if(null==e)return"o";switch(e.charCodeAt(0)){case 91:case 93:case 46:case 34:case 39:return e;case 95:case 36:case 45:return"i";case 9:case 10:case 13:case 160:case 65279:case 8232:case 8233:return"w"}return"i"}function Ae(e){const t=e.trim();return("0"!==e.charAt(0)||!isNaN(parseInt(e)))&&(n=t,Ie.test(n)?function(e){const t=e.charCodeAt(0);return t!==e.charCodeAt(e.length-1)||34!==t&&39!==t?e:e.slice(1,-1)}(t):"*"+t);var n}const ke=new Map;function ye(e,t){return b(e)?e[t]:null}function Ce(e,t){if(!b(e))return null;let n=ke.get(t);if(n||(n=function(e){const t=[];let n,r,o,s,c,a,l,i=-1,u=0,f=0;const _=[];function E(){const t=e[i+1];if(5===u&&"'"===t||6===u&&'"'===t)return i++,o="\\"+t,_[0](),!0}for(_[0]=()=>{void 0===r?r=o:r+=o},_[1]=()=>{void 0!==r&&(t.push(r),r=void 0)},_[2]=()=>{_[0](),f++},_[3]=()=>{if(f>0)f--,u=4,_[0]();else{if(f=0,void 0===r)return!1;if(r=Ae(r),!1===r)return!1;_[1]()}};null!==u;)if(i++,n=e[i],"\\"!==n||!E()){if(s=Oe(n),l=he[u],c=l[s]||l.l||8,8===c)return;if(u=c[0],void 0!==c[1]&&(a=_[c[1]],a&&(o=n,!1===a())))return;if(7===u)return t}}(t),n&&ke.set(t,n)),!n)return null;const r=n.length;let o=e,s=0;for(;s<r;){const e=o[n[s]];if(void 0===e)return null;if(k(o))return null;o=e,s++}return o}const be=e=>e,ge=e=>"",Pe="text",Se=e=>0===e.length?"":function(e,t=""){return e.reduce(((e,n,r)=>0===r?e+n:e+t+n),"")}(e),De=e=>null==e?"":A(e)||D(e)&&e.toString===P?JSON.stringify(e,null,2):String(e);function Ue(e,t){return e=Math.abs(e),2===t?e?e>1?1:0:1:e?Math.min(e,2):0}function ve(e={}){const t=e.locale,n=function(e){const t=E(e.pluralIndex)?e.pluralIndex:-1;return e.named&&(E(e.named.count)||E(e.named.n))?E(e.named.count)?e.named.count:E(e.named.n)?e.named.n:t:t}(e),r=b(e.pluralRules)&&y(t)&&k(e.pluralRules[t])?e.pluralRules[t]:Ue,o=b(e.pluralRules)&&y(t)&&k(e.pluralRules[t])?Ue:void 0,s=e.list||[],c=e.named||{};E(e.pluralIndex)&&function(e,t){t.count||(t.count=e),t.n||(t.n=e)}(n,c);function a(t){const n=k(e.messages)?e.messages(t):!!b(e.messages)&&e.messages[t];return n||(e.parent?e.parent.message(t):ge)}const l=D(e.processor)&&k(e.processor.normalize)?e.processor.normalize:Se,i=D(e.processor)&&k(e.processor.interpolate)?e.processor.interpolate:De,u={list:e=>s[e],named:e=>c[e],plural:e=>e[r(n,e.length,o)],linked:(t,...n)=>{const[r,o]=n;let s="text",c="";1===n.length?b(r)?(c=r.modifier||c,s=r.type||s):y(r)&&(c=r||c):2===n.length&&(y(r)&&(c=r||c),y(o)&&(s=o||s));const l=a(t)(u),i="vnode"===s&&A(l)&&c?l[0]:l;return c?(f=c,e.modifiers?e.modifiers[f]:be)(i,s):i;var f},message:a,type:D(e.processor)&&y(e.processor.type)?e.processor.type:Pe,interpolate:i,normalize:l,values:N({},s,c)};return u}let Re=null;function Me(e){Re=e}const we=xe("function:translate");function xe(e){return t=>Re&&Re.emit(e,t)}const Fe=G.__EXTEND_POINT__,Ye=U(Fe),Xe={NOT_FOUND_KEY:Fe,FALLBACK_TO_TRANSLATE:Ye(),CANNOT_FORMAT_NUMBER:Ye(),FALLBACK_TO_NUMBER_FORMAT:Ye(),CANNOT_FORMAT_DATE:Ye(),FALLBACK_TO_DATE_FORMAT:Ye(),EXPERIMENTAL_CUSTOM_MESSAGE_COMPILER:Ye(),__EXTEND_POINT__:Ye()},We=H.__EXTEND_POINT__,$e=U(We),Ge={INVALID_ARGUMENT:We,INVALID_DATE_ARGUMENT:$e(),INVALID_ISO_DATE_ARGUMENT:$e(),NOT_SUPPORT_NON_STRING_MESSAGE:$e(),NOT_SUPPORT_LOCALE_PROMISE_VALUE:$e(),NOT_SUPPORT_LOCALE_ASYNC_FUNCTION:$e(),NOT_SUPPORT_LOCALE_TYPE:$e(),__EXTEND_POINT__:$e()};function Ke(e){return j(e,null,void 0)}function He(e,t){return null!=t.locale?je(t.locale):je(e.locale)}let Ve;function je(e){if(y(e))return e;if(k(e)){if(e.resolvedOnce&&null!=Ve)return Ve;if("Function"===e.constructor.name){const t=e();if(g(t))throw Ke(Ge.NOT_SUPPORT_LOCALE_PROMISE_VALUE);return Ve=t}throw Ke(Ge.NOT_SUPPORT_LOCALE_ASYNC_FUNCTION)}throw Ke(Ge.NOT_SUPPORT_LOCALE_TYPE)}function Be(e,t,n){return[...new Set([n,...A(t)?t:b(t)?Object.keys(t):y(t)?[t]:[n]])]}function ze(e,t,n){const r=y(n)?n:et,o=e;o.__localeChainCache||(o.__localeChainCache=new Map);let s=o.__localeChainCache.get(r);if(!s){s=[];let e=[n];for(;A(e);)e=Je(s,e,t);const c=A(t)||!D(t)?t:t.default?t.default:null;e=y(c)?[c]:c,A(e)&&Je(s,e,!1),o.__localeChainCache.set(r,s)}return s}function Je(e,t,n){let r=!0;for(let o=0;o<t.length&&C(r);o++){const s=t[o];y(s)&&(r=Qe(e,t[o],n))}return r}function Qe(e,t,n){let r;const o=t.split("-");do{r=qe(e,o.join("-"),n),o.splice(-1,1)}while(o.length&&!0===r);return r}function qe(e,t,n){let r=!1;if(!e.includes(t)&&(r=!0,t)){r="!"!==t[t.length-1];const o=t.replace(/!/g,"");e.push(o),(A(n)||D(n))&&n[o]&&(r=n[o])}return r}const Ze=-1,et="en-US",tt="",nt=e=>`${e.charAt(0).toLocaleUpperCase()}${e.substr(1)}`;let rt,ot,st;function ct(e){rt=e}function at(e){ot=e}function lt(e){st=e}let it=null;const ut=e=>{it=e},ft=()=>it;let _t=null;const Et=e=>{_t=e},dt=()=>_t;let pt=0;function mt(e={}){const t=k(e.onWarn)?e.onWarn:v,n=y(e.version)?e.version:"9.13.1",r=y(e.locale)||k(e.locale)?e.locale:et,o=k(r)?et:r,s=A(e.fallbackLocale)||D(e.fallbackLocale)||y(e.fallbackLocale)||!1===e.fallbackLocale?e.fallbackLocale:o,c=D(e.messages)?e.messages:{[o]:{}},a=D(e.datetimeFormats)?e.datetimeFormats:{[o]:{}},l=D(e.numberFormats)?e.numberFormats:{[o]:{}},i=N({},e.modifiers||{},{upper:(e,t)=>"text"===t&&y(e)?e.toUpperCase():"vnode"===t&&b(e)&&"__v_isVNode"in e?e.children.toUpperCase():e,lower:(e,t)=>"text"===t&&y(e)?e.toLowerCase():"vnode"===t&&b(e)&&"__v_isVNode"in e?e.children.toLowerCase():e,capitalize:(e,t)=>"text"===t&&y(e)?nt(e):"vnode"===t&&b(e)&&"__v_isVNode"in e?nt(e.children):e}),u=e.pluralRules||{},f=k(e.missing)?e.missing:null,_=!C(e.missingWarn)&&!p(e.missingWarn)||e.missingWarn,E=!C(e.fallbackWarn)&&!p(e.fallbackWarn)||e.fallbackWarn,d=!!e.fallbackFormat,m=!!e.unresolving,L=k(e.postTranslation)?e.postTranslation:null,T=D(e.processor)?e.processor:null,h=!C(e.warnHtmlMessage)||e.warnHtmlMessage,I=!!e.escapeParameter,O=k(e.messageCompiler)?e.messageCompiler:rt,g=k(e.messageResolver)?e.messageResolver:ot||ye,P=k(e.localeFallbacker)?e.localeFallbacker:st||Be,S=b(e.fallbackContext)?e.fallbackContext:void 0,U=e,R=b(U.__datetimeFormatters)?U.__datetimeFormatters:new Map,M=b(U.__numberFormatters)?U.__numberFormatters:new Map,w=b(U.__meta)?U.__meta:{};pt++;const x={version:n,cid:pt,locale:r,fallbackLocale:s,messages:c,modifiers:i,pluralRules:u,missing:f,missingWarn:_,fallbackWarn:E,fallbackFormat:d,unresolving:m,postTranslation:L,processor:T,warnHtmlMessage:h,escapeParameter:I,messageCompiler:O,messageResolver:g,localeFallbacker:P,fallbackContext:S,onWarn:t,__meta:w};return x.datetimeFormats=a,x.numberFormats=l,x.__datetimeFormatters=R,x.__numberFormatters=M,__INTLIFY_PROD_DEVTOOLS__&&function(e,t,n){Re&&Re.emit("i18n:init",{timestamp:Date.now(),i18n:e,version:t,meta:n})}(x,n,w),x}function Nt(e,t,n,r,o){const{missing:s,onWarn:c}=e;if(null!==s){const r=s(e,n,t,o);return y(r)?r:t}return t}function Lt(e,t,n){e.__localeChainCache=new Map,e.localeFallbacker(e,n,t)}function Tt(e,t){const n=t.indexOf(e);if(-1===n)return!1;for(let s=n+1;s<t.length;s++)if(r=e,o=t[s],r!==o&&r.split("-")[0]===o.split("-")[0])return!0;var r,o;return!1}function ht(e){return t=>function(e,t){const n=t.b||t.body;if(1===(n.t||n.type)){const t=n,r=t.c||t.cases;return e.plural(r.reduce(((t,n)=>[...t,It(e,n)]),[]))}return It(e,n)}(t,e)}function It(e,t){const n=t.s||t.static;if(n)return"text"===e.type?n:e.normalize([n]);{const n=(t.i||t.items).reduce(((t,n)=>[...t,Ot(e,n)]),[]);return e.normalize(n)}}function Ot(e,t){const n=t.t||t.type;switch(n){case 3:{const e=t;return e.v||e.value}case 9:{const e=t;return e.v||e.value}case 4:{const n=t;return e.interpolate(e.named(n.k||n.key))}case 5:{const n=t;return e.interpolate(e.list(null!=n.i?n.i:n.index))}case 6:{const n=t,r=n.m||n.modifier;return e.linked(Ot(e,n.k||n.key),r?Ot(e,r):void 0,e.type)}case 7:{const e=t;return e.v||e.value}case 8:{const e=t;return e.v||e.value}default:throw new Error(`unhandled node type on format message part: ${n}`)}}const At=e=>e;let kt=Object.create(null);const yt=e=>b(e)&&(0===e.t||0===e.type)&&("b"in e||"body"in e);function Ct(e,t={}){let n=!1;const r=t.onError||B;return t.onError=e=>{n=!0,r(e)},l(a({},Te(e,t)),{detectError:n})}const bt=(e,t)=>{if(!y(e))throw Ke(Ge.NOT_SUPPORT_NON_STRING_MESSAGE);{!C(t.warnHtmlMessage)||t.warnHtmlMessage;const n=(t.onCacheKey||At)(e),r=kt[n];if(r)return r;const{code:o,detectError:s}=Ct(e,t),c=new Function(`return ${o}`)();return s?c:kt[n]=c}};function gt(e,t){if(__INTLIFY_JIT_COMPILATION__&&!__INTLIFY_DROP_MESSAGE_COMPILER__&&y(e)){!C(t.warnHtmlMessage)||t.warnHtmlMessage;const n=(t.onCacheKey||At)(e),r=kt[n];if(r)return r;const{ast:o,detectError:s}=Ct(e,l(a({},t),{location:!1,jit:!0})),c=ht(o);return s?c:kt[n]=c}{const t=e.cacheKey;if(t){const n=kt[t];return n||(kt[t]=ht(e))}return ht(e)}}const Pt=()=>"",St=e=>k(e);function Dt(e,...t){const{fallbackFormat:n,postTranslation:r,unresolving:o,messageCompiler:s,fallbackLocale:c,messages:a}=e,[l,i]=Rt(...t),u=C(i.missingWarn)?i.missingWarn:e.missingWarn,f=C(i.fallbackWarn)?i.fallbackWarn:e.fallbackWarn,_=C(i.escapeParameter)?i.escapeParameter:e.escapeParameter,d=!!i.resolvedMessage,p=y(i.default)||C(i.default)?C(i.default)?s?l:()=>l:i.default:n?s?l:()=>l:"",m=n||""!==p,L=He(e,i);_&&function(e){A(e.list)?e.list=e.list.map((e=>y(e)?h(e):e)):b(e.named)&&Object.keys(e.named).forEach((t=>{y(e.named[t])&&(e.named[t]=h(e.named[t]))}))}(i);let[T,I,O]=d?[l,L,a[L]||{}]:Ut(e,l,L,c,f,u),k=T,g=l;if(d||y(k)||yt(k)||St(k)||m&&(k=p,g=k),!(d||(y(k)||yt(k)||St(k))&&y(I)))return o?-1:l;let P=!1;const S=St(k)?k:vt(e,l,I,k,g,(()=>{P=!0}));if(P)return k;const D=function(e,t,n,r){const{modifiers:o,pluralRules:s,messageResolver:c,fallbackLocale:a,fallbackWarn:l,missingWarn:i,fallbackContext:u}=e,f=r=>{let o=c(n,r);if(null==o&&u){const[,,e]=Ut(u,r,t,a,l,i);o=c(e,r)}if(y(o)||yt(o)){let n=!1;const s=vt(e,r,t,o,r,(()=>{n=!0}));return n?Pt:s}return St(o)?o:Pt},_={locale:t,modifiers:o,pluralRules:s,messages:f};e.processor&&(_.processor=e.processor);r.list&&(_.list=r.list);r.named&&(_.named=r.named);E(r.plural)&&(_.pluralIndex=r.plural);return _}(e,I,O,i),U=function(e,t,n){const r=t(n);return r}(0,S,ve(D)),v=r?r(U,l):U;if(__INTLIFY_PROD_DEVTOOLS__){const t={timestamp:Date.now(),key:y(l)?l:St(k)?k.key:"",locale:I||(St(k)?k.locale:""),format:y(k)?k:St(k)?k.source:"",message:v};t.meta=N({},e.__meta,ft()||{}),we(t)}return v}function Ut(e,t,n,r,o,s){const{messages:c,onWarn:a,messageResolver:l,localeFallbacker:i}=e,u=i(e,r,n);let f,_={},E=null;for(let d=0;d<u.length&&(f=u[d],_=c[f]||{},null===(E=l(_,t))&&(E=_[t]),!(y(E)||yt(E)||St(E)));d++)if(!Tt(f,u)){const n=Nt(e,t,f,0,"translate");n!==t&&(E=n)}return[E,f,_]}function vt(e,t,n,r,o,s){const{messageCompiler:c,warnHtmlMessage:a}=e;if(St(r)){const e=r;return e.locale=e.locale||n,e.key=e.key||t,e}if(null==c){const e=()=>r;return e.locale=n,e.key=t,e}const l=c(r,function(e,t,n,r,o,s){return{locale:t,key:n,warnHtmlMessage:o,onError:e=>{throw s&&s(e),e},onCacheKey:e=>f(t,n,e)}}(0,n,o,0,a,s));return l.locale=n,l.key=t,l.source=r,l}function Rt(...e){const[t,n,r]=e,o={};if(!(y(t)||E(t)||St(t)||yt(t)))throw Ke(Ge.INVALID_ARGUMENT);const s=E(t)?String(t):(St(t),t);return E(n)?o.plural=n:y(n)?o.default=n:D(n)&&!m(n)?o.named=n:A(n)&&(o.list=n),E(r)?o.plural=r:y(r)?o.default=r:D(r)&&N(o,r),[s,o]}function Mt(e,...t){const{datetimeFormats:n,unresolving:r,fallbackLocale:o,onWarn:s,localeFallbacker:c}=e,{__datetimeFormatters:a}=e,[l,i,u,f]=xt(...t);C(u.missingWarn)?u.missingWarn:e.missingWarn;C(u.fallbackWarn)?u.fallbackWarn:e.fallbackWarn;const _=!!u.part,E=He(e,u),d=c(e,o,E);if(!y(l)||""===l)return new Intl.DateTimeFormat(E,f).format(i);let p,L={},T=null;for(let m=0;m<d.length&&(p=d[m],L=n[p]||{},T=L[l],!D(T));m++)Nt(e,l,p,0,"datetime format");if(!D(T)||!y(p))return r?-1:l;let h=`${p}__${l}`;m(f)||(h=`${h}__${JSON.stringify(f)}`);let I=a.get(h);return I||(I=new Intl.DateTimeFormat(p,N({},T,f)),a.set(h,I)),_?I.formatToParts(i):I.format(i)}const wt=["localeMatcher","weekday","era","year","month","day","hour","minute","second","timeZoneName","formatMatcher","hour12","timeZone","dateStyle","timeStyle","calendar","dayPeriod","numberingSystem","hourCycle","fractionalSecondDigits"];function xt(...e){const[t,n,r,o]=e,s={};let c,a={};if(y(t)){const e=t.match(/(\d{4}-\d{2}-\d{2})(T|\s)?(.*)/);if(!e)throw Ke(Ge.INVALID_ISO_DATE_ARGUMENT);const n=e[3]?e[3].trim().startsWith("T")?`${e[1].trim()}${e[3].trim()}`:`${e[1].trim()}T${e[3].trim()}`:e[1].trim();c=new Date(n);try{c.toISOString()}catch(l){throw Ke(Ge.INVALID_ISO_DATE_ARGUMENT)}}else if(d(t)){if(isNaN(t.getTime()))throw Ke(Ge.INVALID_DATE_ARGUMENT);c=t}else{if(!E(t))throw Ke(Ge.INVALID_ARGUMENT);c=t}return y(n)?s.key=n:D(n)&&Object.keys(n).forEach((e=>{wt.includes(e)?a[e]=n[e]:s[e]=n[e]})),y(r)?s.locale=r:D(r)&&(a=r),D(o)&&(a=o),[s.key||"",c,s,a]}function Ft(e,t,n){const r=e;for(const o in n){const e=`${t}__${o}`;r.__datetimeFormatters.has(e)&&r.__datetimeFormatters.delete(e)}}function Yt(e,...t){const{numberFormats:n,unresolving:r,fallbackLocale:o,onWarn:s,localeFallbacker:c}=e,{__numberFormatters:a}=e,[l,i,u,f]=Wt(...t);C(u.missingWarn)?u.missingWarn:e.missingWarn;C(u.fallbackWarn)?u.fallbackWarn:e.fallbackWarn;const _=!!u.part,E=He(e,u),d=c(e,o,E);if(!y(l)||""===l)return new Intl.NumberFormat(E,f).format(i);let p,L={},T=null;for(let m=0;m<d.length&&(p=d[m],L=n[p]||{},T=L[l],!D(T));m++)Nt(e,l,p,0,"number format");if(!D(T)||!y(p))return r?-1:l;let h=`${p}__${l}`;m(f)||(h=`${h}__${JSON.stringify(f)}`);let I=a.get(h);return I||(I=new Intl.NumberFormat(p,N({},T,f)),a.set(h,I)),_?I.formatToParts(i):I.format(i)}const Xt=["localeMatcher","style","currency","currencyDisplay","currencySign","useGrouping","minimumIntegerDigits","minimumFractionDigits","maximumFractionDigits","minimumSignificantDigits","maximumSignificantDigits","compactDisplay","notation","signDisplay","unit","unitDisplay","roundingMode","roundingPriority","roundingIncrement","trailingZeroDisplay"];function Wt(...e){const[t,n,r,o]=e,s={};let c={};if(!E(t))throw Ke(Ge.INVALID_ARGUMENT);const a=t;return y(n)?s.key=n:D(n)&&Object.keys(n).forEach((e=>{Xt.includes(e)?c[e]=n[e]:s[e]=n[e]})),y(r)?s.locale=r:D(r)&&(c=r),D(o)&&(c=o),[s.key||"",a,s,c]}function $t(e,t,n){const r=e;for(const o in n){const e=`${t}__${o}`;r.__numberFormatters.has(e)&&r.__numberFormatters.delete(e)}}"boolean"!=typeof __INTLIFY_PROD_DEVTOOLS__&&(T().__INTLIFY_PROD_DEVTOOLS__=!1),"boolean"!=typeof __INTLIFY_JIT_COMPILATION__&&(T().__INTLIFY_JIT_COMPILATION__=!1),"boolean"!=typeof __INTLIFY_DROP_MESSAGE_COMPILER__&&(T().__INTLIFY_DROP_MESSAGE_COMPILER__=!1);export{Rt as A,Dt as B,Ge as C,et as D,xt as E,Mt as F,Wt as G,Yt as H,yt as I,St as J,ze as K,ct as L,tt as M,Xt as N,at as O,lt as P,Xe as Q,Et as R,gt as S,bt as T,Ce as U,N as a,y as b,b as c,C as d,j as e,A as f,T as g,D as h,E as i,p as j,k,i as l,u as m,wt as n,m as o,M as p,O as q,mt as r,Me as s,Ft as t,Lt as u,$t as v,U as w,ut as x,dt as y,Ze as z};