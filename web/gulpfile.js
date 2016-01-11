var gulp = require('gulp');
var $ = require('gulp-load-plugins')({lazy: true});
var source = require('vinyl-source-stream');
var browserify = require('browserify');
var watchify = require('watchify');
var reactify = require('reactify');

gulp.task('styles', function () {
    log('Compiling SASS ==> CSS');
    return gulp
        .src('./sass/**/*.scss')
        .pipe($.plumber({ errorHandler: handleError }))
        .pipe($.sourcemaps.init())
        .pipe($.sass())
        .pipe($.sourcemaps.write())
        .pipe($.autoprefixer({browsers: ['last 2 version', '> 5%']}))
        .pipe(gulp.dest('./public/css/'))
        .pipe($.livereload());
});


gulp.task('sass-watcher', function () {
    $.livereload.listen();
    gulp.watch('./sass/**/*.scss', ['styles']);
});

gulp.task('react', function(){
    log('Compiling React ==> JS');
    var bundler = browserify({
        entries: ['./jsx/main.jsx'],
        transform: [reactify],
        extensions: ['.jsx'],
        debug: true,
        cache: {},
        packageCache: {},
        fullPaths: true
    });

    function build(file){
        if(file) log('recompiling ' + file);

        return bundler
            .bundle()
            .on('error', $.util.log.bind($.util, 'Browserify error'))
            .pipe(source('app.js'))
            .pipe(gulp.dest('./public/js'))
            .pipe($.livereload());
    };

    return build();
});

gulp.task('jsx-watcher', function(){
    $.livereload.listen();
    var bundler = watchify(browserify({
        entries: ['./jsx/main.jsx'],
        transform: [reactify],
        extensions: ['.jsx'],
        debug: true,
        cache: {},
        packageCache: {},
        fullPaths: true
    }));

    function build(file){
        if(file) log('recompiling ' + file);

        return bundler
            .bundle()
            .on('error', $.util.log.bind($.util, 'Browserify error'))
            .pipe(source('app.js'))
            .pipe(gulp.dest('./public/js'))
            .pipe($.livereload());
    };

    build();
    bundler.on('update', build);
});

gulp.task('default', ['styles', 'sass-watcher', 'jsx-watcher']);


//////////////////////////////////////////////////////////////////////


function log (msg) {
    if (typeof(msg) === 'object') {
        for (var item in msg) {
            if (msg.hasOwnProperty(item)) {
                $.util.log($.util.colors.yellow(msg[item]));
            }
        }
    } else {
        $.util.log($.util.colors.yellow(msg));
    }
}

function handleError (err) {
    $.util.beep();
    log(err);
};
