package main

import (
	"context"
)

type Ptime struct{}

func (m *Ptime) Test(ctx context.Context) error {

	pgSvc, err := dag.Container().From("postgres:11").WithEnvVariable("POSTGRES_PASSWORD", "postgres").WithExposedPort(5432).AsService().Start(ctx)
	if err != nil {
		return err
	}

	memcSvc, err := dag.Container().From("memcached").WithExposedPort(11211).AsService().Start(ctx)
	if err != nil {
		return err
	}

	_, err = dag.Container().From("ruby:2.7").
		WithServiceBinding("postgres", pgSvc).
		WithServiceBinding("memcached", memcSvc).
		WithEnvVariable("RAILS_TEST_DB_NAME", "postgres").
		WithEnvVariable("RAILS_TEST_DB_USERNAME", "postgres").
		WithEnvVariable("RAILS_TEST_DB_PASSWORD", "postgres").
		WithEnvVariable("RAILS_ENV", "test").
		WithEnvVariable("CI", "true").
		WithEnvVariable("PGDATESTYLE", "German").
		WithExec([]string{"apt-get", "-yqq", "install", "libpq-dev"}).
		WithExec([]string{"gem", "install", "bundler", "--version", "'~> 2'"}).
		WithExec([]string{"bundle", "install", "--jobs", "4", "--retry", "3"}).
		WithExec([]string{"bundle", "exec", "rails", "db:create"}).
		WithExec([]string{"bundle", "exec", "rails", "db:migrate"}).
		WithExec([]string{"bundle", "exec", "rails", "test"}).Sync(ctx)

	pgSvc.Stop(ctx)
	memcSvc.Stop(ctx)

	return err
}

func (m *Ptime) Lint(ctx context.Context) error {
	return nil
}

func (m *Ptime) Build(ctx context.Context) error {
	return nil
}
