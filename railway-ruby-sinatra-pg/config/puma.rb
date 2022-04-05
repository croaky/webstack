# Increasing the number of workers will increase resting memory.
# Increasing the number of threads will increase potential bloat
# when processes are responding to heavy requests.
#
# A low number of workers and threads provides adequate
# performance for most apps, even under load,
# while maintaining a low risk of overusing memory.
workers ENV.fetch("PUMA_WORKERS", 2).to_i
threads_count = ENV.fetch("PUMA_THREADS", 2).to_i
threads(threads_count, threads_count)

preload_app!

port ENV.fetch("PORT", 9292)
environment ENV.fetch("APP_ENV", "dev")
