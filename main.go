package main

import (
	"context"
)

type Ptime struct{}

func (m *Ptime) Test(ctx context.Context) error {

	workDir := dag.Host().Directory(".")

	pgSvc := dag.Container().From("postgres:11").WithEnvVariable("POSTGRES_PASSWORD", "postgres").WithExposedPort(5432).AsService()

	memcSvc := dag.Container().From("memcached").WithExposedPort(11211).AsService()

	_, err := dag.Container().From("ruby:3.2.1").
		WithServiceBinding("postgres", pgSvc).
		WithServiceBinding("memcached", memcSvc).
		WithEnvVariable("RAILS_TEST_DB_NAME", "postgres").
		WithEnvVariable("RAILS_TEST_DB_USERNAME", "postgres").
		WithEnvVariable("RAILS_TEST_DB_PASSWORD", "postgres").
		WithEnvVariable("RAILS_TEST_DB_HOST", "postgres").
		WithEnvVariable("RAILS_ENV", "test").
		WithEnvVariable("CI", "true").
		WithEnvVariable("PGDATESTYLE", "German").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "-yqq", "install", "libpq-dev", "nodejs", "npm", "rubygems", "libvips-dev"}).
		WithDirectory("/src", workDir).
		WithWorkdir("/src").
		WithExec([]string{"mv", "_vendor", "vendor"}).
		WithExec([]string{"wget", "https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb"}).
		WithExec([]string{"sh", "-c", "dpkg -i google-chrome-stable_current_amd64.deb || apt -yqq --fix-broken install"}).
		WithExec([]string{"gem", "install", "bundler", "--version", "~> 2"}).
		WithExec([]string{"bundle", "install", "--jobs", "4", "--retry", "3"}).
		WithExec([]string{"bundle", "exec", "rails", "db:create"}).
		WithExec([]string{"bundle", "exec", "rails", "db:migrate"}).
		WithExec([]string{"bundle", "exec", "rails", "assets:precompile"}).
		WithExec([]string{"bundle", "exec", "rails", "test"}).Sync(ctx)

	pgSvc.Stop(ctx)
	memcSvc.Stop(ctx)

	return err
}

/*
func (m *Ptime) Puilt(ctx context.Context) error {

	workdir := dag.Host().Directory(".")

	rubyContainer := dag.Gale().Run(GaleRunOpts{
		Source:   workdir,
		Workflow: "Code Style Review",
	}).Sync()

	rubyContainer.Export(ctx, "./export.tar.gz")

	_, err := rubyContainer.
		WithEnvVariable("RAILS_TEST_DB_NAME", "postgres").
		WithEnvVariable("RAILS_TEST_DB_USERNAME", "postgres").
		WithEnvVariable("RAILS_TEST_DB_PASSWORD", "postgres").
		WithEnvVariable("RAILS_ENV", "test").
		WithEnvVariable("CI", "true").
		WithEnvVariable("PGDATESTYLE", "German").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "-yqq", "install", "libpq-dev"}).
		//WithWorkdir("/src").
		WithExec([]string{"gem", "install", "bundler", "--version", "~> 2"}).
		WithExec([]string{"bundle", "install", "--jobs", "4", "--retry", "3"}).
		WithEnvVariable("CACHE_BUSERT", fmt.Sprintf("%d", time.Now().UnixMilli())).
		WithExec([]string{"bundle", "exec", "rails", "db:create"}).
		WithExec([]string{"bundle", "exec", "rails", "db:migrate"}).
		WithExec([]string{"bundle", "exec", "rails", "test"}).Sync(ctx)
	return err
}

func (m *Ptime) Lint(ctx context.Context) error {
	return nil
}

func (m *Ptime) Build(ctx context.Context) error {
	return nil
}*/

// inputs:
//
//	github_token:
//	  description: 'GITHUB_TOKEN'
//	  default: ${{ github.token }}
//	rubocop_version:
//	  description: 'Rubocop version'
//	rubocop_extensions:
//	  description: 'Rubocop extensions'
//	  default: 'rubocop-rails rubocop-performance rubocop-rspec rubocop-i18n rubocop-rake'
//	rubocop_flags:
//	  description: 'Rubocop flags. (rubocop <rubocop_flags>)'
//	  default: ''
//	tool_name:
//	  description: 'Tool name to use for reviewdog reporter'
//	  default: 'rubocop'
//	level:
//	  description: 'Report level for reviewdog [info,warning,error]'
//	  default: 'error'
//	reporter:
//	  description: |
//	    Reporter of reviewdog command [github-pr-check,github-check,github-pr-review].
//	    Default is github-pr-check.
//	  default: 'github-pr-check'
//	filter_mode:
//	  description: |
//	    Filtering mode for the reviewdog command [added,diff_context,file,nofilter].
//	    Default is added.
//	  default: 'added'
//	fail_on_error:
//	  description: |
//	    Exit code for reviewdog when errors are found [true,false]
//	    Default is `false`.
//	  default: 'false'
//	reviewdog_flags:
//	  description: 'Additional reviewdog flags'
//	  default: ''
//	workdir:
//	  description: "The directory from which to look for and run Rubocop. Default '.'"
//	  default: '.'
//	skip_install:
//	  description: "Do not install Rubocop or its extensions. Default: `false`"
//	  default: 'false'
//	use_bundler:
//	  description: "Run Rubocop with bundle exec. Default: `false`"
//	  default: 'false'
func (m *Ptime) RubocopReviewdog(ctx context.Context) error {

	_, err := dag.Container().From("node:20").
		WithWorkdir("/src").
		WithExec([]string{"wget", "https://raw.githubusercontent.com/reviewdog/action-rubocop/master/script.sh"}).
		WithExec([]string{"chmod", "+x", "./script.sh"}).
		WithExec([]string{"sh", "-c", "./script.sh"}).Sync(ctx)

	return err
}
