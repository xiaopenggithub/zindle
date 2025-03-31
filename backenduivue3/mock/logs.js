export default [
  {
    url: '/api/logs',
    method: 'get',
    response: ({ query }) => {
      const { current = 1, pageSize = 10, searchQuery = '', logLevel = '', dateRange = '' } = query
      
      // Generate 100 items total
      const totalItems = 100
      const allLogs = Array(totalItems).fill(null).map((_, index) => ({
        id: `${index + 1}`,
        time: new Date(Date.now() - index * 3600000).toISOString(),
        level: ['INFO', 'WARNING', 'ERROR', 'DEBUG'][Math.floor(Math.random() * 4)],
        user: ['admin', 'user1', 'user2', 'system'][Math.floor(Math.random() * 4)],
        ip: `192.168.1.${Math.floor(Math.random() * 255)}`,
        action: ['LOGIN', 'LOGOUT', 'CREATE', 'UPDATE', 'DELETE'][Math.floor(Math.random() * 5)],
        details: `操作日志 ${index + 1}`
      }))

      // Filter logs based on search criteria
      let filteredLogs = [...allLogs]
      
      if (searchQuery) {
        filteredLogs = filteredLogs.filter(log => 
          log.details.toLowerCase().includes(searchQuery.toLowerCase()) ||
          log.user.toLowerCase().includes(searchQuery.toLowerCase()) ||
          log.ip.includes(searchQuery)
        )
      }

      if (logLevel) {
        filteredLogs = filteredLogs.filter(log => log.level === logLevel)
      }

      if (dateRange && dateRange !== '[]') {
        try {
          const [startDate, endDate] = JSON.parse(dateRange)
          if (startDate && endDate) {
            filteredLogs = filteredLogs.filter(log => {
              const logDate = new Date(log.time)
              return logDate >= new Date(startDate) && logDate <= new Date(endDate)
            })
          }
        } catch (e) {
          console.error('Date range parsing error:', e)
        }
      }

      // Calculate pagination
      const start = (parseInt(current) - 1) * parseInt(pageSize)
      const end = start + parseInt(pageSize)
      const paginatedLogs = filteredLogs.slice(start, end)

      return {
        code: 200,
        data: {
          list: paginatedLogs,
          total: filteredLogs.length
        }
      }
    }
  },
  {
    url: '/api/logs/:id',
    method: 'delete',
    response: () => ({
      code: 200,
      message: 'Success'
    })
  },
  {
    url: '/api/logs/export',
    method: 'get',
    response: () => ({
      code: 200,
      data: 'Mock export data'
    })
  }
] 