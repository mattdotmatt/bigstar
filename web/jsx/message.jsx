var React = require('react');
var ReactDOM = require('react-dom');
var jQuery = require('jquery');

module.exports = React.createClass({

    fade: function(elem) {
        jQuery(elem).fadeIn("slow");
        jQuery(elem).fadeOut(2000);
    },

    componentDidMount: function() {
        var elem = ReactDOM.findDOMNode(this);
        this.fade(elem);
    },

    componentDidUpdate: function() {
        var elem = ReactDOM.findDOMNode(this);
        this.fade(elem);
    },

    render: function () {
        return <div id="success-alert" className="alert alert-success" role="alert">{this.props.message}</div>
    }
});