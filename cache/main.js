const args = {
  port: 8060,
  max_cache: 64
};

process.argv.slice(2).map((arg) => {
  args[arg.split("=")[0]] = arg.split("=")[1];
});

