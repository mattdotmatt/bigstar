var React = require('react');

module.exports = React.createClass({

    render: function () {
        return <div id="success-alert" className="alert alert-success" role="alert">{this.props.message}</div>
    }
});