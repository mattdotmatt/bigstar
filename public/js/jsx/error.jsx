var React = require('react');

module.exports = React.createClass({

    render: function () {
        return <div className="alert alert-warning" role="alert">{this.props.errorMsg}</div>
    }
});